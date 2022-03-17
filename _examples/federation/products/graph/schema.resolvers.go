package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/00security/gqlgen/_examples/federation/products/graph/generated"
	model1 "github.com/00security/gqlgen/_examples/federation/reviews/graph/model"
)

func (r *manufacturerResolver) Name(ctx context.Context, obj *model1.Manufacturer) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) Upc(ctx context.Context, obj *model1.Product) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) Name(ctx context.Context, obj *model1.Product) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) Price(ctx context.Context, obj *model1.Product) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TopProducts(ctx context.Context, first *int) ([]*model1.Product, error) {
	return hats, nil
}

// Manufacturer returns generated.ManufacturerResolver implementation.
func (r *Resolver) Manufacturer() generated.ManufacturerResolver { return &manufacturerResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type manufacturerResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
