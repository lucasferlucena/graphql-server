package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lucasferlucena/graphql-server/db/models"
	"github.com/lucasferlucena/graphql-server/graph/generated"
)

func (r *queryResolver) Videos(ctx context.Context) ([]*models.Video, error) {
	var videos []*models.Video
	r.DB.Find(&videos)

	return videos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	r.DB.Find(&users)

	return users, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
