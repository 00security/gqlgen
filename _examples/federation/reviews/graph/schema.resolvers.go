package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	model1 "github.com/00security/gqlgen/_examples/federation/accounts/graph/model"
	"github.com/00security/gqlgen/_examples/federation/reviews/graph/generated"
	"github.com/00security/gqlgen/_examples/federation/reviews/graph/model"
)

func (r *reviewResolver) Author(ctx context.Context, obj *model.Review) (*model1.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Reviews(ctx context.Context, obj *model1.User) ([]*model.Review, error) {
	var productReviews []*model.Review
	for _, review := range reviews {
		if review.Author.ID == obj.ID {
			productReviews = append(productReviews, review)
		}
	}
	return productReviews, nil
}

// Review returns generated.ReviewResolver implementation.
func (r *Resolver) Review() generated.ReviewResolver { return &reviewResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type reviewResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
