// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Invoice struct {
	Invno      string    `json:"invno"`
	CustType   string    `json:"cust_type"`
	Custphone  string    `json:"custphone"`
	Invdate    time.Time `json:"invdate"`
	Invoicedat string    `json:"invoicedat"`
	Custno     string    `json:"custno"`
	Discount   int       `json:"discount"`
	Discamount float64   `json:"discamount"`
	BalAc      float64   `json:"bal_ac"`
	Balance    float64   `json:"balance"`
	Taxamount  float64   `json:"taxamount"`
	Subtotal   float64   `json:"subtotal"`
	Totamount  float64   `json:"totamount"`
	Totbev     float64   `json:"totbev"`
	Paidm      int       `json:"paidm"`
	Timeset    string    `json:"timeset"`
	Otime      string    `json:"otime"`
	Lunch      string    `json:"lunch"`
	Server     string    `json:"server"`
	TableNo    string    `json:"table_no"`
	NoPerson   string    `json:"no_person"`
	Delivery   string    `json:"delivery"`
	Mark       string    `json:"mark"`
	Checkprint string    `json:"checkprint"`
	Orderprint int       `json:"orderprint"`
	Tip        float64   `json:"tip"`
	CashPaid   float64   `json:"cash_paid"`
	ChangeAmt  float64   `json:"change_amt"`
	Stateno    string    `json:"stateno"`
	Finished   string    `json:"finished"`
	Unsend     string    `json:"unsend"`
	ServiceCh  float64   `json:"service_ch"`
	Upstatus   int       `json:"upstatus"`
	Discrateam float64   `json:"discrateam"`
	Nontaxsale string    `json:"nontaxsale"`
	Paidst     string    `json:"paidst"`
	Vipno      string    `json:"vipno"`
	Discbywho  string    `json:"discbywho"`
	Custname   string    `json:"custname"`
	Void       int       `json:"void"`
	Split      int       `json:"split"`
	Deltime    string    `json:"deltime"`
	Lineprt    string    `json:"lineprt"`
	Batch      int       `json:"batch"`
	Paytime    string    `json:"paytime"`
	Onlines    int       `json:"onlines"`
	Cooked     int       `json:"cooked"`
	Mapmark    string    `json:"mapmark"`
	Note       string    `json:"note"`
}

type SalesInput struct {
	Start   *time.Time `json:"start"`
	End     *time.Time `json:"end"`
	RangeBy *string    `json:"range_by"`
	Type    *string    `json:"type"`
}

type SalesOverTime struct {
	TimePoint    int       `json:"time_point"`
	CompleteDate time.Time `json:"complete_date"`
	TotalAmount  float64   `json:"total_amount"`
}

type SalesStats struct {
	Today                          float64 `json:"today"`
	TodayYesterdayDiff             float64 `json:"today_yesterday_diff"`
	ThisYearTodayLastYearTodayDiff float64 `json:"this_year_today_last_year_today_diff"`
	Total                          float64 `json:"total"`
	Profit                         float64 `json:"profit"`
}
