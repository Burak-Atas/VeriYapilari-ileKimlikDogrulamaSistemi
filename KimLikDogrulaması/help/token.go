package help

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreatedNewToken(email, isim, soyisim, id string) string {
	claims := &GirisDetayları{
		Email:      email,
		First_name: isim,
		Last_name:  soyisim,
		Uid:        id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	return token
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

type GirisDetayları struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string

	jwt.StandardClaims
}
