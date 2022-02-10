package main

import (
	"fmt"
	"net/http"
	"time"
)

type JSONAPIResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"` // timestamp of the challenge load (ISO format yyyy-MM-dd'T'HH:mm:ssZZ)
	Hostname    string    `json:"hostname"`     // the hostname of the site where the reCAPTCHA was solved
	ErrorCodes  []int     `json:"error-codes"`  //optional
}

func RecaptchaForm(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
                  <html lang="en">
                  <head>
                   <meta charset="UTF-8">
                   <title>Golang reCAPTCHA Signup Form</title>
                   <script src='https://www.google.com/recaptcha/api.js'></script>
                  </head>
                  <body>
                  <h1>Golang reCAPTCHA Signup Form</h1>
                  <form method="POST" action="/signup">
                   Username : <input type="text" name="username">
                   <br>
                   Password : <input type="password" name="password">
                   <br>
                   <div class="g-recaptcha" data-sitekey="6LcG3ckaAAAAALuyrJgq4RWTuKXaMP3sX4Ypyrx6"></div>
                   <br>
                   <input type="submit" value="Submit">
                  </form>
                  </body>
                  </html>`

	w.Write([]byte(html))
}

func Signup(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		fmt.Println("username : ", r.FormValue("username"))
		fmt.Println("password : ", r.FormValue("password"))

		response := r.FormValue("g-recaptcha-response")
		fmt.Println("g-recaptcha-response : ", response)

		// did we get a proper recaptcha response? if null, redirect back to sigup page
		if response == "" {
			// user press submit button without passing reCAPTCHA test
			// abort
			http.Redirect(w, r, "/", 301)
			return // return control to stop execution, otherwise it will continue
		}

		// this is server side validation example
		// it is optional, but no harm adding this on your signup verification process.

		// if you are running this example on localhost
		// get your own IP address from https://whatismyipaddress.com/
		// and assign it to the variable ip

		// get end user's IP address
		//remoteip, _, _ := net.SplitHostPort(r.RemoteAddr) fmt.Println("remote ip : ", remoteip)

		// to verify if the recaptcha is REAL. we must send
		// secret + response + remoteip(optional) to postURL

		//secret := "6Leip58aAAAAAAzabSUTJJhvgpIA0t9Qw7tF1J5K"
		//postURL := "https://www.google.com/recaptcha/api/siteverify"

		//postStr := url.Values{"secret": {secret}, "response": {response}}

		//responsePost, err := http.PostForm(postURL, postStr)
		//
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//
		//defer responsePost.Body.Close()
		//body, err := ioutil.ReadAll(responsePost.Body)
		//
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//
		//// this part is for server side verification
		//var APIResp JSONAPIResponse
		//
		//json.Unmarshal(body, &APIResp)
		//fmt.Println(APIResp)

		// see https://www.socketloop.com/tutorials/golang-output-or-print-out-json-stream-encoded-data

		// w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`<!DOCTYPE html>
		<html lang="en">
		<head>
		 <meta charset="UTF-8">
		 <title>Golang reCAPTCHA Signup Form</title>
		</head>
		<body>
		<h1>Golang reCAPTCHA Signup Form</h1>
		<h6>Captcha: ` + response + `</h6>
		<a href="http://localhost:5000/">Back</a>
		</body>
		</html>`))

		//once everything is verified, you can proceed to save the user data into your database
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", RecaptchaForm)
	mux.HandleFunc("/signup", Signup)

	http.ListenAndServe(":5000", mux)
}
