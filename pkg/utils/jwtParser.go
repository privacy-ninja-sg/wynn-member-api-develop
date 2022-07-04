package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	Expires int64  `json:"exp"`
	Uid     int    `json:"uid"`
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Picture string `json:"picture"`
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx, jwtSecret string) (*TokenMetadata, error) {
	token, err := verifyToken(c, jwtSecret)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// Expires time.
		expires := int64(claims["exp"].(float64))
		uid := int(claims["uid"].(float64))
		uuid := claims["uuid"].(string)
		name := claims["name"].(string)
		status := claims["status"].(string)
		picture := claims["picture"].(string)

		return &TokenMetadata{
			Expires: expires,
			Uid:     uid,
			Name:    name,
			UUID:    uuid,
			Status:  status,
			Picture: picture,
		}, nil
	}

	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx, jwtSecret string) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
