package table

import (
	"github.com/guregu/null"
)

// MessageLog message_log table
type MessageLog struct {
	JobID               null.String `field:"job_id" property:"primary_key"`
	MessageType         null.Int    `field:"message_type"`
	Message             []byte      `field:"message"`
	CreateDateTimeStamp null.Time   `field:"create_datetime_stamp"`
}

// TableName message_log name
func (*MessageLog) TableName() string {
	return "message_log"
}
