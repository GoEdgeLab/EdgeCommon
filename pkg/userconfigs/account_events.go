// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package userconfigs

// 所有账户相关的事件类型

type AccountEventType = string

const (
	AccountEventTypeCharge   AccountEventType = "charge"   // 充值
	AccountEventTypeAward    AccountEventType = "award"    // 赠送
	AccountEventTypeBuyPlan  AccountEventType = "buyPlan"  // 购买套餐
	AccountEventTypePayBill  AccountEventType = "payBill"  // 支付账单
	AccountEventTypeRefund   AccountEventType = "refund"   // 退款
	AccountEventTypeWithdraw AccountEventType = "withdraw" // 提现
)

type AccountEvent struct {
	Name        string           `json:"name"`        // 名称
	Code        AccountEventType `json:"code"`        // 代号
	Description string           `json:"description"` // 描述
	IsPositive  bool             `json:"isPositive"`  // 是否为正向
}

var AccountIncomeEventTypes = []AccountEventType{AccountEventTypeCharge}    // 收入
var AccountExpenseEventTypes = []AccountEventType{AccountEventTypeWithdraw} // 支出

// FindAllAccountEventTypes 查找所有的事件类型
func FindAllAccountEventTypes() []*AccountEvent {
	return []*AccountEvent{
		{
			Name:        "充值",
			Code:        AccountEventTypeCharge,
			Description: "为用户账户充值。",
			IsPositive:  true,
		},
		{
			Name:        "赠送",
			Code:        AccountEventTypeAward,
			Description: "为用户账户赠送余额。",
			IsPositive:  true,
		},
		{
			Name:        "购买套餐",
			Code:        AccountEventTypeBuyPlan,
			Description: "购买套餐支出。",
			IsPositive:  false,
		},
		{
			Name:        "支付账单",
			Code:        AccountEventTypePayBill,
			Description: "支付账单支出。",
			IsPositive:  false,
		},
		{
			Name:        "退款",
			Code:        AccountEventTypeRefund,
			Description: "退款到用户账户。",
			IsPositive:  true,
		},
		{
			Name:        "提现",
			Code:        AccountEventTypeWithdraw,
			Description: "用户从账户提现。",
			IsPositive:  false,
		},
	}
}

// FindAccountEvent 根据事件类型查找事件定义
func FindAccountEvent(eventType AccountEventType) *AccountEvent {
	for _, e := range FindAllAccountEventTypes() {
		if e.Code == eventType {
			return e
		}
	}
	return nil
}
