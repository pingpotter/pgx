package table

import "github.com/guregu/null"

// Product product table
type Product struct {
	ProductName         null.String `field:"product_name" property:"primary_key"`
	ProductGroup        null.String `field:"product_group"`
	ProductType         null.String `field:"product_type"`
	ProductSubType      null.String `field:"product_subtype"`
	IsRevolving         null.Bool   `field:"is_revolving"`
	TransactionPlanName null.String `field:"transaction_plan_name"`
	InterestPlanName    null.String `field:"interest_plan_name"`
	PaymentPlanName     null.String `field:"payment_plan_name"`
	PenaltyPlanName     null.String `field:"penalty_plan_name"`
	GracePlanName       null.String `field:"grace_plan_name"`
	Description         null.String `field:"description"`
}

// TableName product name
func (*Product) TableName() string {
	return "product"
}
