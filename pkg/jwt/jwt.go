package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

// JWTClaims represents the claims in the JWT
type JWTClaims struct {
	AadharNumber string `json:"aadharNumber"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	jwt.RegisteredClaims
}

// JWTManager handles JWT operations
type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

// NewJWTManager creates a new JWTManager
func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:     secretKey,
		tokenDuration: tokenDuration,
	}
}

// Generate creates a new JWT token for a user
func (m *JWTManager) Generate(aadharNumber, email, role string) (string, error) {
	claims := JWTClaims{
		AadharNumber: aadharNumber,
		Email:        email,
		Role:         role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secretKey))
}

// Verify validates the token string and returns the claims
func (m *JWTManager) Verify(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(m.secretKey), nil
		},
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// ExtractTokenFromHeader extracts the JWT token from the Authorization header
func ExtractTokenFromHeader(c *fiber.Ctx) string {
	bearerToken := c.Get("Authorization")
	if len(bearerToken) > 7 && bearerToken[:7] == "Bearer " {
		return bearerToken[7:]
	}
	return ""
}

// Middleware creates a Fiber middleware for JWT authentication
func (m *JWTManager) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := ExtractTokenFromHeader(c)
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}

		claims, err := m.Verify(tokenString)
		if err != nil {
			if errors.Is(err, ErrExpiredToken) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "token expired",
				})
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		// Store user information in the context for later use
		c.Locals("aadharNumber", claims.AadharNumber)
		c.Locals("email", claims.Email)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}

// Optional middleware that verifies JWT if present but doesn't require it
func (m *JWTManager) OptionalAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := ExtractTokenFromHeader(c)
		if tokenString != "" {
			claims, err := m.Verify(tokenString)
			if err == nil {
				c.Locals("aadharNumber", claims.AadharNumber)
				c.Locals("email", claims.Email)
				c.Locals("role", claims.Role)
				c.Locals("authenticated", true)
			}
		}
		return c.Next()
	}
}

// RoleMiddleware creates a middleware that checks if the user has the required role
func (m *JWTManager) RoleMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// First apply the JWT middleware
		err := m.Middleware()(c)
		if err != nil {
			return err
		}

		// Then check the role
		role := c.Locals("role").(string)
		if role != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "forbidden access",
			})
		}

		return c.Next()
	}
}
