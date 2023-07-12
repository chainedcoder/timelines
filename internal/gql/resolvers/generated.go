package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// import (
// 	"chainedcoder/timelines/internal/gql/models"
// 	"context"
// )

// func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
//     panic("not implemented")
// }
// func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
//     panic("not implemented")
// }
// func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
//     panic("not implemented")
// }

// func (r *queryResolver) Users(ctx context.Context, userID *string) ([]*models.User, error) {
//     records := []*models.User{
//         &models.User{
//             ID:     new(string),
//             Email:  new(string),
//             UserID: new(string),
//         },
//     }
//     *records[0].ID = "ec17af15-e354-440c-a09f-69715fc8b595"
//     *records[0].Email = "your@email.com"
//     *records[0].UserID = "UserID-1"
//     return records, nil
// }