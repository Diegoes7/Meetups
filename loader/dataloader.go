package loader

import (
	"context"
	"net/http"
	"time"

	"github.com/Diegoes7/meetups/models"
	"github.com/go-pg/pg/v10"
)

// ContextKey is a custom type to avoid collisions with other context keys
type ContextKey string

const userLoaderKey ContextKey = "userLoader"

// DataLoaderMiddleware sets up a UserLoader and injects it into the request context
func DataLoaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := &UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*models.User, []error) {
				var users []*models.User

				//$ Query users based on IDs and put them in allocated memory
				err := db.Model(&users).Where("id IN (?)", pg.In(ids)).Select()
				if err != nil {
					return nil, []error{err}
				}

				//* create a map with a specific length  
				u := make(map[string]*models.User, len(users))

				//! loop over the map and set key/value pair
				for _, user := range users {
					u[user.ID] = user
				}

				//! key/value pair the object is assigned to a key in a struct like dictionary
				// {
				// 	"1":  User_1,
				// 	"2": USer_2,
				// }

				//& create a slice, not a map
				result := make([]*models.User, len(ids))

				for i, id := range ids {
					result[i] = u[id] //! this put user {} in result[{user with specific ID}, {user with specific ID}, 
					//!... till the end of len of the ids] so put user in un array and load it with same order every time
				}

				return result, nil
			},
		}

		// Add the loader to the context
		ctx := context.WithValue(r.Context(), userLoaderKey, userLoader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserLoader(ctx context.Context) *UserLoader {
	loader, ok := ctx.Value(userLoaderKey).(*UserLoader)
	if !ok {
		panic("UserLoader not found in context") // Or return an error if you prefer
	}
	return loader
}
