package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type contextKey struct{}

type User struct {
	ID    int
	Roles []string
}

/*
func AddUserToContext(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, contextKey{}, user)
}

func GetUserFromContext(ctx context.Context) (User, bool) {
	user, ok := ctx.Value(contextKey{}).(User)
	return user, ok
}
*/

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		raw := strings.TrimPrefix(
			c.GetHeader("Authorization"),
			"Bearer ",
		)

		claims, err := ParseToken(raw, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		user := User{
			ID:    claims.UserID,
			Roles: claims.Roles,
		}

		c.Set("user", user)
		c.Next()
	}
}
