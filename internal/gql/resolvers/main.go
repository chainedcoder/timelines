package resolvers

import (
	"context"

	"chainedcoder/timelines/internal/logger"

	"chainedcoder/timelines/pkg/utils"

	"chainedcoder/timelines/internal/gql"
	"chainedcoder/timelines/internal/orm"
	dbm "chainedcoder/timelines/internal/orm/models"
)

// Resolver is a modifable struct that can be used to pass on properties used
// in the resolvers, such as DB access
type Resolver struct {
	ORM *orm.ORM
}

// Mutation exposes mutation methods
func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}

// Query exposes query methods
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func getCurrentUser(ctx context.Context) *dbm.User {
	cu := ctx.Value(utils.ProjectContextKeys.UserCtxKey).(*dbm.User)
	logger.Debugf("currentUser: %s - %s", cu.Email, cu.ID)
	return cu
}