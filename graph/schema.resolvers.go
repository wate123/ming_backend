package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ming_backend/graph/generated"
	"ming_backend/graph/model"
)

func (r *queryResolver) Invoice(ctx context.Context) ([]*model.Invoice, error) {
	var invoices []*model.Invoice
	r.DB.Find(&invoices)
	return invoices, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Invoices(ctx context.Context) ([]*model.Invoice, error) {
	var invoices []*model.Invoice
	r.DB.Find(&invoices)
	return invoices, nil
}

type mutationResolver struct{ *Resolver }
