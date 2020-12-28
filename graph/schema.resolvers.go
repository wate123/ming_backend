package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ming_backend/graph/generated"
	"ming_backend/graph/model"
	"time"
)

func (r *queryResolver) Invoices(ctx context.Context) ([]*model.Invoice, error) {
	var invoices []*model.Invoice
	r.DB.Find(&invoices)
	return invoices, nil
}

//TODO Profit
//
func (r *queryResolver) GetAllSalesStats(ctx context.Context) (*model.SalesStats, error) {
	var saleStats model.SalesStats
	var yesterday float64
	var lastYearToday float64
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	r.DB.Table("invoice").Select("SUM(totamount)").Where("invdate = ?",
		today).Scan(&saleStats.Today)
	r.DB.Table("invoice").Select("SUM(totamount)").Where("invdate = ?",
		today.AddDate(0, 0, -1)).Scan(&yesterday)

	r.DB.Table("invoice").Select("SUM(totamount)").Where("invdate = ?",
		today.AddDate(-1, 0, 0)).Scan(&lastYearToday)
	r.DB.Table("invoice").Select("SUM(totamount)").Where("invdate BETWEEN ? AND ?",
		time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC), now).Scan(&saleStats.Total)
	saleStats.TodayYesterdayDiff = saleStats.Today - yesterday
	saleStats.ThisYearTodayLastYearTodayDiff = saleStats.Today - lastYearToday
	return &saleStats, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
