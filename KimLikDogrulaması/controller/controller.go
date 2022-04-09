package controller

import (
	"KimLikDogrulaması/help"
	linkedlist "KimLikDogrulaması/linkedList"
	"KimLikDogrulaması/models"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

var Dogrulama = validator.New()
var linker linkedlist.LinkedList

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "sign up sayfasında"})
		var kullanici models.Kullanici
		DogrulamaHatasi := Dogrulama.Struct(kullanici)
		if DogrulamaHatasi != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": DogrulamaHatasi})
			return
		}

		count := linker.SearchEmail(kullanici.Email)
		if count != 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "eposta sistemde zaten kayitli"})
		}
		sifre := HashPasword(kullanici.Password)
		kullanici.Password = sifre
		kullanici.Kimlik = ObjectID()
		token := help.CreatedNewToken(kullanici.Email, kullanici.İsim, kullanici.Soyisim, kullanici.Kimlik)

		kullanici.Token = token
		fmt.Println("\n\n")
		fmt.Println(kullanici.Email)
		fmt.Println(kullanici.Token)
		fmt.Println(linker.GetSize())
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

func ObjectID() string {
	karakterler := "abcdefghijklmnoprstuvxwyz"
	karakterler += strings.ToUpper(karakterler)
	karakterler += "1234567890+%&/"
	var newID string = ""
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().Unix())
		x := rand.Intn(len(karakterler))
		newID = newID + string(karakterler[x])
	}
	return newID
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
