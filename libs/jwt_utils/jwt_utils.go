package jwt_utils

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type TokenInfo struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Admin  bool      `json:"admin"`
	jwt.RegisteredClaims
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return getSecret(), nil
}

func GetTokenInfo(token string) *TokenInfo {
	token_obj, _ := jwt.ParseWithClaims(token, &TokenInfo{}, keyFunc)
	token_info, ok := token_obj.Claims.(*TokenInfo)

	if ok && token_obj.Valid {
		return token_info
	} else {
		return &TokenInfo{}
	}
}

func GetTokenInfoFromRequest(r *http.Request) *TokenInfo {
	return GetTokenInfo(GetTokenString(r))
}

func GenerateToken(user_id uuid.UUID, email string, admin bool) string {
	token_info := TokenInfo{
		UserID: user_id,
		Email:  email,
		Admin:  admin,
	}

	token_obj := jwt.NewWithClaims(jwt.SigningMethodHS256, token_info)
	token, err := token_obj.SignedString(getSecret())

	if err != nil {
		fmt.Println(err)
	}
	return token
}

func GetTokenString(r *http.Request) string {
	return r.Header.Get("Authorization")
}

func getSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}
