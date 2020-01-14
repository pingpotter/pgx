package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/guregu/null"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type server struct {
	pool *pgxpool.Pool
	tx   pgx.Tx
	postgresRepository
}

func main() {

	log.Println("start api")

	s := &server{}
	s.Start(context.Background())
	defer s.Stop()

	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *server) Start(ctx context.Context) error {
	if s.pool != nil {
		return nil
	}

	url := "postgres://postgres:password@pg:5432/dloan_account?connect_timeout=10&pool_max_conns=10"

	pool, err := connectWithRetries(ctx, url)
	if err != nil {
		return err
	}
	s.pool, s.postgresRepository.qi = pool, pool

	return nil
}

func (s *server) Begin(ctx context.Context) error {

	x1, err := s.qi.Begin(ctx)
	s.qix = x1

	return err
}

// func (s *server) Rollback(ctx context.Context) error {
// 	return s.Rollback(ctx context.Context)
// }

func (s *server) Stop() {
	if s.pool == nil {
		return
	}
	s.pool.Close()
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		const sql = `SELECT true`
		var tmp bool

		if err := s.qi.QueryRow(context.Background(), sql).Scan(&tmp); err != nil {
			log.Println("error: ", err)
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))

	case "POST":

		// opt := pgx.TxOptions{
		// 	// IsoLevel: pgx.Serializable,
		// 	IsoLevel: pgx.ReadCommitted,
		// 	// IsoLevel: pgx.RepeatableRead,
		// }

		var prod null.String
		if err := s.qi.QueryRow(context.Background(), "select product_name from product where product_name=$1", "7200120090001").Scan(&prod); err != nil {
			log.Println("error: ", err)
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		err := s.Begin(context.Background())
		// tx, err := s.pool.BeginTx(context.Background(), opt)
		// tx, err := s.Db.Begin(context.Background())
		if err != nil {
			log.Println("error: ", err)
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		defer func() {
			err := s.qix.Rollback(context.Background())
			log.Println("this defer:", err)
		}()

		var seqNumber null.Int

		if err := s.qix.QueryRow(context.Background(), "select account_number from config_account where key_value=$1", "init-account-number").Scan(&seqNumber); err != nil {
			log.Println("error: ", err)
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		// panic("thispanic")

		startNumber := seqNumber.Int64 + 1
		accountNumber, accountInit, _ := calculateDigit(startNumber)
		log.Println("accountNumber:", accountNumber)

		_, err = s.qix.Exec(context.Background(), "update config_account set account_number=$1", null.IntFrom(accountInit))
		if err != nil {
			log.Println("error: ", err)
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		_, err = s.qix.Exec(context.Background(), "insert into account(account_number) values($1)", null.IntFrom(accountNumber))
		if err != nil {
			log.Println("error: ", err)
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		_, err = s.qix.Exec(context.Background(), "insert into message_log(job_id) values($1)", strconv.FormatInt(accountNumber, 10))
		if err != nil {
			log.Println("error: ", err)
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		uuid := null.StringFrom(strconv.FormatInt(time.Now().UnixNano(), 10) + strconv.FormatInt(accountNumber, 10))
		_, err = s.qix.Exec(context.Background(), "insert into opened_account_journal(uid,account_number) values($1,$2)", uuid, null.IntFrom(accountNumber))
		if err != nil {
			log.Println("error: ", err)
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		if err := s.qix.Commit(context.Background()); err != nil {
			log.Println("error: ", err)
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}

}

// calculateDigit : Condition for generate loan account
func calculateDigit(input int64) (int64, int64, error) {

	accountInput := strconv.FormatInt(input, 10)

	var remaining, beginingDigit, lastDigit int

	for i := 1; i < 20; i++ {
		account := accountInput
		remaining = findRemaining(accountInput)
		if remaining == 1 {
			accountInput = strconv.FormatInt(input+1, 10)
			continue
		} else if remaining == 0 {
			beginingDigit = splitString(account, 1, 11)
			lastDigit = remaining
			break
		} else {
			beginingDigit = splitString(account, 1, 11)
			lastDigit = 11 - remaining
			break
		}
	}

	outputBeginingDigit := strconv.Itoa(beginingDigit)
	outputLastDigit := strconv.Itoa(lastDigit)
	loanAccountNumber, err := strconv.ParseInt(outputBeginingDigit+outputLastDigit, 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return loanAccountNumber, int64(beginingDigit), nil
}

func findRemaining(loanAccountNumber string) int {

	calNumber := 6
	digit := 0

	for i := 1; i < 12; i++ {
		if i == 6 {
			calNumber = 7
		}
		digit += (splitString(loanAccountNumber, i, i) * calNumber)
		calNumber--

	}
	return digit % 11
}

// SplitString : Split String
func splitString(input string, start int, end int) int {
	runes := []rune(input)
	if len(runes) >= end {
		output, _ := strconv.Atoi(string(runes[start-1 : end]))
		return output
	}
	output, _ := strconv.Atoi(input)
	return output
}
