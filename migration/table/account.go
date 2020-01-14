package table

import (
	"github.com/guregu/null"
)

// Account account table
type Account struct {
	AccountNumber       null.Int    `field:"account_number" property:"primary_key"`
	ProductName         null.String `field:"product_name"`
	CustomerNumber      null.String `field:"customer_number"`
	CustomerType        null.String `field:"customer_type"`
	AccountName         null.String `field:"account_name"`
	AccountBranch       null.Int    `field:"account_branch"`
	ResponseUnit        null.Int    `field:"response_unit"`
	CreditTermNumber    null.Int    `field:"credit_term_number"`
	CreditTermUnit      null.String `field:"credit_term_unit"`
	DisbursementAccount null.String `field:"disbursement_account"`
	OpenAccountJobID    null.String `field:"open_account_jobid"`
	OpenDate            null.Time   `field:"open_date"`
	OpenDatetimeStamp   null.Time   `field:"open_datetime_stamp"`
	ApplicationID       null.String `field:"application_id"`
	ClosedDate          null.Time   `field:"closed_date"`
	MaturityDate        null.Time   `field:"maturity_date"`
}

// TableName account name
func (*Account) TableName() string {
	return "account"
}
