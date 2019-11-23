package web

import (
	"time"
)

// Account info
// AccountName:账户名, Amount:金额, BindedAt:参与时间,
// Amount:参与的金额, TotalNumbers:累计参与的次数，
// TotalAmounts:累计参与的金额
type Accounts struct {
	ID            uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`
	AccountAddr   string    `gorm:"column:account_addr;type:char(64);not null;" json:"account_addr" form:"account_addr"`
	BindedTime    time.Time `gorm:"column:binded_time;type:datetime;not null;" json:"binded_time" form:"binded_time"`
	Amount        int       `gorm:"column:amount;type:int;not null;" json:"amount" form:"amount"`
}

type UserInfos struct {
	ID            uint64  `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`
	AccountAddr   string  `gorm:"column:account_addr;type:char(64);unique_index:uk_account_addr;not null;" json:"account_addr" form:"account_addr"`
	TotalNumbers  int     `gorm:"column:total_numbers;type:int;not null;" json:"total_numbers" form:"total_numbers"`
	TotalAmounts  int     `gorm:"column:total_amounts;type:int;not null;" json:"total_amounts" form:"total_amounts"`
}

type UserSettle struct {
	ID            uint64     `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`
	AccountAddr   string     `gorm:"column:account_addr;type:char(64);not null;" json:"account_addr" form:"account_addr"`
	BindedAmount  int        `gorm:"column:amount;type:int;not null;" json:"amount" form:"amount"`
	SettleTime    time.Time  `gorm:"column:settle_time;type:datetime;not null;" json:"settle_time" form:"settle_time"`
	SettleAmount  int        `gorm:"column:settle_amount;type:int;not null;" json:"settle_amount" form:"settle_amount"`
}

