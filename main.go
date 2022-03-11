package main

import (
	"auth/captcha"
	configs "auth/config"
	services "auth/google"
	"auth/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	// Initialize Viper across the applib  cation
	configs.InitializeViper()

	// Initialize Logger across the application

	// Initialize Oauth2 Services
	services.InitializeOAuthGoogle()

	// Routes for the application
	http.HandleFunc("/", handler.HandleMain)
	http.HandleFunc("/login-gl", services.HandleGoogleLogin)
	http.HandleFunc("/login-captcha", captcha.RecaptchaForm)
	http.HandleFunc("/signup", captcha.Signup)
	http.HandleFunc("/callback-gl", services.CallBackFromGoogle)

	fmt.Println("Server Started @ 7001...")
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), nil))
}

//func CreateToken(userid uint64) (string, error) {
//	var err error
//	//Creating Access Token
//	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
//	atClaims := jwt.MapClaims{}
//	atClaims["authorized"] = true
//	atClaims["user_id"] = userid
//	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
//	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
//	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
//	if err != nil {
//		return "", err
//	}
//	return token, nil
//}
//
