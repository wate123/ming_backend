package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/snabb/isoweek"
	"log"
	"ming_backend/graph/generated"
	"ming_backend/graph/model"
	"ming_backend/util"
	"time"
)

func (r *queryResolver) Invoices(ctx context.Context) ([]*model.Invoice, error) {
	var invoices []*model.Invoice
	r.DB.Find(&invoices)
	return invoices, nil
}

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

func (r *queryResolver) GetSalesByDate(ctx context.Context, input model.DateInput) ([]*model.SalesOverTime, error) {
	var salesOverTime []*model.SalesOverTime
	timeRange := util.RangeBy(*input.RangeBy)
	baseQuery := r.DB.Table("invoice").Select(fmt.Sprintf("%s AS time_point, SUM(totamount) AS total_amount", timeRange))
	if input.Start == nil && input.End == nil {
		log.Fatal("No start date or end date are given")
	} else if input.Start != nil && input.End != nil {
		baseQuery.Where("invdate BETWEEN ? AND ?", input.Start, input.End).Group(timeRange).Find(&salesOverTime)
	} else if input.Start != nil {
		baseQuery.Where("invdate >= ?", input.Start).Group(timeRange).Find(&salesOverTime)
	} else {
		baseQuery.Where("invdate <= ?", input.End).Group(timeRange).Find(&salesOverTime)
	}
	return salesOverTime, nil
}

/*
	return all the sales by month, from the first day of this year up to today
 */
func (r *queryResolver) GetSalesFromThisYear(ctx context.Context) ([]*model.SalesOverTime, error) {
	t := time.Now()
	firstDayThisYear := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	rangeBy := "month"
	return r.GetSalesByDate(ctx, model.DateInput{Start: &firstDayThisYear, RangeBy: &rangeBy})
}

/*
	return all the sales by day, from the first day of this month up to today
 */
func (r *queryResolver) GetSalesFromThisMonth(ctx context.Context) ([]*model.SalesOverTime, error) {
	t := time.Now()
	firstDayThisMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	rangeBy := "day"
	return r.GetSalesByDate(ctx, model.DateInput{Start: &firstDayThisMonth, RangeBy: &rangeBy})
}

/*
	return all the sales by day, starting from the first day of this week
 */
func (r *queryResolver) GetSalesFromThisWeek(ctx context.Context) ([]*model.SalesOverTime, error) {
	t := time.Now()
	_, thisWeek := isoweek.FromDate(t.Year(), t.Month(), t.Day())
	firstDayThisWeek := isoweek.StartTime(t.Year(), thisWeek, time.UTC)
	rangeBy := "day"
	return r.GetSalesByDate(ctx, model.DateInput{Start: &firstDayThisWeek, RangeBy: &rangeBy})
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
