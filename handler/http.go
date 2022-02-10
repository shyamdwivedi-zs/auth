package handler

import (
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
	"strings"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(IndexPage))
}

func HandleLogin(w http.ResponseWriter, r *http.Request, oauthConf *oauth2.Config, oauthStateString string) {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Auth path-->",URL.String())
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	parameters.Add("access_type", "offline")
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	fmt.Println("url is--->",url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}