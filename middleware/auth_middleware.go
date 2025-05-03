package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/Diegoes7/meetups/models"
	"github.com/Diegoes7/meetups/postgres"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
	"github.com/pkg/errors"
)

const CurrentUserKey = "currentUser"

func AuthMiddleware(repo postgres.UserRepo) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := parseToken(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims) //* cast it to other type
			if !ok || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			userID, ok := claims["jti"].(string)
			if !ok || userID == "" {
				next.ServeHTTP(w, r)
				return
			}

			user, err := repo.GetUserByID(userID)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), CurrentUserKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromToken,
}

func stripBearerPrefixFromToken(token string) (string, error) {
	bearer := "Bearer"
	if len(token) > len(bearer) && strings.EqualFold(token[0:len(bearer)], bearer) {
		return token[len(bearer)+1:], nil
	}
	return token, nil
}

var authExtractor = request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

var tokenStr string

func parseToken(r *http.Request) (*jwt.Token, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}

	// Try cookie first
	if cookie, err := r.Cookie("authToken"); err == nil {
		tokenStr = cookie.Value
	} else {
		// Fallback to headers/query param
		var err error
		tokenStr, err = authExtractor.ExtractToken(r)
		if err != nil {
			return nil, errors.Wrap(err, "extractToken error")
		}
	}

	token, err := request.ParseFromRequest(r, authExtractor, keyFunc)
	if err != nil {
		return nil, errors.Wrap(err, "parseToken error")
	}

	return token, nil
}

func GetCurrentUserFromCTX(ctx context.Context) (*models.User, error) {
	user, ok := ctx.Value(CurrentUserKey).(*models.User)
	if !ok || user == nil || user.ID == "" {
		return nil, errors.New("no valid user in context")
	}
	return user, nil
}
