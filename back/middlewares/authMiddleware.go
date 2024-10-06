package middleware

import (
	"back/config"
	"back/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func AuthMiddleware(allowedRoles ...uint8) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		prefix := "Bearer "

		authHeader = strings.TrimPrefix(authHeader, prefix)

		token, err := jwt.ParseWithClaims(authHeader, &services.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return config.JWTSecret, nil
		})

		if err != nil {
			fmt.Printf("Token parsing error: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if claims, ok := token.Claims.(*services.Claims); ok && token.Valid {
			// Vérifie si l'utilisateur a un des rôles autorisés
			for _, role := range allowedRoles {
				if config.HasRequiredRole(claims.Roles, role) {
					c.Set("userID", claims.UserID)
					c.Next()
					return
				}
			}

			// Si aucun rôle n'a été trouvé
			fmt.Printf("Insufficient permissions: User roles: %d, Required roles: %v\n", claims.Roles, allowedRoles)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
		} else {
			fmt.Printf("Token validation error: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}

		/*if claims, ok := token.Claims.(*services.Claims); ok && token.Valid {
			fmt.Printf("User roles: %d, Required roles: %d, Has required role: %v\n", claims.Roles, requiredRole, config.HasRequiredRole(claims.Roles, requiredRole))
			if !config.HasRequiredRole(claims.Roles, requiredRole) {
				fmt.Printf("Insufficient permissions: User roles: %d, Required roles: %d\n", claims.Roles, requiredRole)
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
				return
			}

			c.Set("userID", claims.UserID)
			c.Next()
		} else {
			fmt.Printf("Token validation error: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}*/
	}
}
