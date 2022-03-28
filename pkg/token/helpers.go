package token

import (
	"log"
	"time"

	"github.com/swooosh13/quest-auth/internal/config"

	"github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Login string
	Uid   string
	jwt.StandardClaims
}

func GenerateAllTokens(login, uid string) (signedToken string, signedRefreshToken string, err error) {
	var SECRET_KEY string = config.GetConfig().SecretKey

	claims := &SignedDetails{
		Login: login,
		Uid:   uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}
