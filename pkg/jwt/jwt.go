package jwt

import (
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 24 * 365

var mySercet = []byte("Gaojiayuan")

type MyClaims struct {
	Username string `json:"username"`
	UserID   int64  `json:"userid"`
	jwt.StandardClaims
}

func GenToken(Userid int64, username string) (string, error) {
	claims := MyClaims{
		username,
		Userid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("jwt") * 24)).Unix(),
			Issuer:    "bluebell",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySercet)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySercet, nil
	})

	if token.Valid {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, err
	}
	return nil, err
}
