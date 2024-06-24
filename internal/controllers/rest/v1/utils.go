package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/LevinStrike/lux-backend/internal/core/apperror"
	"github.com/LevinStrike/lux-backend/internal/core/entities"
	"github.com/golang-jwt/jwt"
)

func JsonDecode(buf io.Reader, body any) error {
	if buf == nil {
		return apperror.ErrUnexpectedNillValues
	}
	if err := json.NewDecoder(buf).Decode(body); err != nil {
		return apperror.NewBadRequestError(err)
	}
	return nil
}

var secretKey = []byte("secret-key")

type jwtClaim struct {
	ID  int   `json:"id"`
	EXP int64 `json:"exp"`
	jwt.StandardClaims
}

func createToken(user entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwtClaim{
			ID:  user.GetID(),
			EXP: time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) (int, error) {
	claims := jwtClaim{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	if time.Now().Unix() > claims.EXP {
		return 0, errors.New("token expired")
	}

	return claims.ID, nil
}
