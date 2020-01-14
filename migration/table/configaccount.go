package table

import (
	"github.com/guregu/null"
)

// ConfigAccount config_account table
type ConfigAccount struct {
	KeyValue      null.String `field:"key_value" property:"primary_key"`
	AccountNumber null.Int    `field:"account_number"`
}

// TableName config_account name
func (*ConfigAccount) TableName() string {
	return "config_account"
}
