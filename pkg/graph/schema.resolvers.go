package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/maxsei/ts-label/pkg/graph/generated"
	"github.com/maxsei/ts-label/pkg/graph/model"
)

func (r *mutationResolver) UpdateLabel(ctx context.Context, ids []string, xAxisLabel string, name string, value int) (*model.Label, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetChartIds(ctx context.Context) ([][]string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetChartData(ctx context.Context, ids []string) (*model.ChartData, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
