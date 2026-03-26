package jwt_utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type TokenInfo struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return getSecret(), nil
}

func GetTokenInfo(token string) *TokenInfo {
	token_obj, _ := jwt.ParseWithClaims(token, &TokenInfo{}, keyFunc)
	if token_obj == nil {
		return nil
	}
	token_info, ok := token_obj.Claims.(*TokenInfo)
	if ok && token_obj.Valid {
		return token_info
	}
	return nil
}

func GetTokenInfoFromRequest(r *http.Request) *TokenInfo {
	return GetTokenInfo(GetTokenString(r))
}

func GenerateToken(user_id uuid.UUID, email string) string {
	token_info := TokenInfo{
		UserID: user_id,
		Email:  email,
	}

	token_obj := jwt.NewWithClaims(jwt.SigningMethodHS256, token_info)
	token, err := token_obj.SignedString(getSecret())

	if err != nil {
		fmt.Println(err)
	}
	return token
}

func GetTokenString(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	return strings.TrimPrefix(auth, "Bearer ")
}

func getSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}
