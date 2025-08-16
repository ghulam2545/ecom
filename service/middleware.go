package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const CtxClaimsKey = "claims"

// JWTAuthMiddleware verifies Authorization: Bearer <token> and stores claims in context
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" || !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			return
		}
		tok := strings.TrimPrefix(h, "Bearer ")
		claims, err := ParseToken(tok)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.Set(CtxClaimsKey, claims)
		c.Next()
	}
}

// RequireRoles enforces all of the given roles (strict). Useful when a route must be exactly a role.
func RequireRoles(roles ...string) gin.HandlerFunc {
	roleSet := map[string]struct{}{}
	for _, r := range roles {
		roleSet[r] = struct{}{}
	}
	return func(c *gin.Context) {
		val, exists := c.Get(CtxClaimsKey)
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "no claims"})
			return
		}
		claims := val.(*Claims)
		if _, ok := roleSet[claims.Role]; !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient role"})
			return
		}
		c.Next()
	}
}

// RequireAnyRole allows access if user has any role from the list
func RequireAnyRole(roles ...string) gin.HandlerFunc {
	roleSet := map[string]struct{}{}
	for _, r := range roles {
		roleSet[r] = struct{}{}
	}
	return func(c *gin.Context) {
		val, exists := c.Get(CtxClaimsKey)
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "no claims"})
			return
		}
		claims := val.(*Claims)
		if _, ok := roleSet[claims.Role]; !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient role"})
			return
		}
		c.Next()
	}
}

// GetClaims Helper getters
func GetClaims(c *gin.Context) *Claims {
	val, _ := c.Get(CtxClaimsKey)
	if val == nil {
		return nil
	}
	return val.(*Claims)
}
