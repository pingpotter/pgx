package table

import (
	"github.com/guregu/null"
)

// OpenedAccountJournal opened_account_journal table
type OpenedAccountJournal struct {
	UID           null.String `field:"uid" property:"primary_key"`
	AccountNumber null.Int    `field:"account_number" property:"primary_key"`
	IsSent        null.Bool   `field:"is_sent"`
	CreatedAt     null.Time   `field:"create_datetime_stamp"`
	UpdatedAt     null.Time   `field:"update_datetime_stamp"`
}

// TableName opened_account_journal name
func (*OpenedAccountJournal) TableName() string {
	return "opened_account_journal"
}
