package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"ming_backend/graph/generated"
	"ming_backend/graph/model"
	"ming_backend/util"
	"time"

	"github.com/snabb/isoweek"
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

func (r *queryResolver) GetSalesByDate(ctx context.Context, input model.SalesInput) ([]*model.SalesOverTime, error) {
	var salesOverTime []*model.SalesOverTime
	timeRange := util.RangeBy(*input.RangeBy)
	orderType := "%%"
	if input.Type != nil {
		orderType = fmt.Sprintf("%%%s%%", *input.Type) // escape %, format orderType as %pattern%
	}
	query := r.DB.Table("invoice").Select(
		fmt.Sprintf("%s AS time_point, invdate AS complete_date, SUM(totamount) AS total_amount", timeRange))
	if input.Start == nil && input.End == nil {
		log.Fatal("No start date or end date are given")
	} else if input.Start != nil && input.End != nil {
		query = query.Where("custphone LIKE ? AND invdate BETWEEN ? AND ?", orderType, input.Start, input.End)
	} else if input.Start != nil {
		query = query.Where("custphone LIKE ? AND invdate >= ?", orderType, input.Start)
	} else {
		query = query.Where("custphone LIKE ? AND invdate <= ?", orderType, input.End)
	}
	query.Group(timeRange).Order("invdate").Find(&salesOverTime)
	return salesOverTime, nil
}

func (r *queryResolver) GetSalesFromThisYear(ctx context.Context, input *string) ([]*model.SalesOverTime, error) {
	t := time.Now()
	firstDayThisYear := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	rangeBy := "month"
	return r.GetSalesByDate(ctx, model.SalesInput{Start: &firstDayThisYear, RangeBy: &rangeBy, Type: input})
}

func (r *queryResolver) GetSalesFromThisMonth(ctx context.Context, input *string) ([]*model.SalesOverTime, error) {
	t := time.Now()
	firstDayThisMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	rangeBy := "day"
	return r.GetSalesByDate(ctx, model.SalesInput{Start: &firstDayThisMonth, RangeBy: &rangeBy, Type: input})
}

func (r *queryResolver) GetSalesFromThisWeek(ctx context.Context, input *string) ([]*model.SalesOverTime, error) {
	t := time.Now()
	_, thisWeek := isoweek.FromDate(t.Year(), t.Month(), t.Day())
	firstDayThisWeek := isoweek.StartTime(t.Year(), thisWeek, time.UTC)
	rangeBy := "day"
	return r.GetSalesByDate(ctx, model.SalesInput{Start: &firstDayThisWeek, RangeBy: &rangeBy, Type: input})
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
