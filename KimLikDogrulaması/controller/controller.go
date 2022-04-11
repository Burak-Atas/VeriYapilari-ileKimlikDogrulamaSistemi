package controller

import (
	"KimLikDogrulaması/help"
	linkedlist "KimLikDogrulaması/linkedList"
	"KimLikDogrulaması/models"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

var Dogrulama = validator.New()
var linker linkedlist.LinkedList

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		count := 0

		var kullanici models.Kullanici
		DogrulamaHatasi := Dogrulama.Struct(kullanici)
		if DogrulamaHatasi != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": DogrulamaHatasi})
			return
		}
		fmt.Println(count)

		count = linker.SearchEmail(kullanici.Email)
		if count != 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "eposta sistemde zaten kayitli"})
			return
		}
		sifre := HashPasword(kullanici.Password)
		kullanici.Password = sifre
		kullanici.Kimlik = string(GenerateZoneID())
		token := help.CreatedNewToken(kullanici.Email, kullanici.İsim, kullanici.Soyisim, kullanici.Kimlik)

		kullanici.Token = token

		fmt.Println("kullanici tokeni -->>>" + kullanici.Token)

		linker.InsertLast(kullanici.Kimlik, kullanici.İsim, kullanici.Soyisim, kullanici.Email, kullanici.Password)
		c.JSON(http.StatusOK, gin.H{"message": "kullanici başarıyla eklenmiştir"})

	}
}

func HashPasword(pasword string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pasword), 9)
	if err != nil {
		log.Fatal(err)
	}
	/*
		password := "alsk1234"	gibi bir şifrenindönen değeri

		$2a$14$x7IjjKJoV0CQmO1YhlLmm.uwDzs29dMcSrUJdZsr1gc/XViUwecYi
	*/
	return string(bytes)
}

var (
	JwtLength   = 16
	ZoneIDBegin = []byte{'T', 'O', 'K', 'E', 'N', 'J', 'W', 'T'}
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateZoneID() []byte {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, JwtLength)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return append(ZoneIDBegin, b...)
}

func LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var kullanici models.Kullanici

		var bulunanKullanici models.Kullanici

		bulundu := linker.SearchEmail(kullanici.Email)
		if bulundu == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "kullanici bulunamadı"})
			return
		}
		bulunanKullanici.Email = kullanici.Email

		saglandi := SifreDogrulama(bulunanKullanici.Password, kullanici.Password)
		if !saglandi {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "sifreler uyusmadı"})
			return
		}

		if bulunanKullanici.Email == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		x := linker.SearchEmail(bulunanKullanici.Email)
		if x != 1 {
			log.Println("kullanici email hatası")
		}
		c.JSON(http.StatusOK, gin.H{"message": kullanici.Email,
			"":       kullanici.Kimlik,
			"error":  kullanici.İsim,
			"succes": kullanici.Soyisim,
		})

	}
}

func SifreDogrulama(sifre1, sifre2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(sifre1), []byte(sifre2))
	saglandi := true
	if err != nil {
		log.Println("sifre uyusmadı")
		saglandi = false
	}
	return saglandi
}

func GetUSers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "kullanici listesi",
		})
		if linker.GetSize() == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "kullanici bulunamadı"})
			return
		}
		for i := 0; i < linker.GetSize(); i++ {
			c.JSON(http.StatusOK, gin.H{
				"message": linker.GetItems()[i],
			})
		}

	}
}
