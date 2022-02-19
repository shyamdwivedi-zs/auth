package google

import (
	"auth/handler"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"text/template"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfGl = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:7000/callback-gl",
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/calendar",
			"https://www.googleapis.com/auth/drive.file", "https://www.googleapis.com/auth/drive", "https://mail.google.com/",
			"https://www.googleapis.com/auth/gmail.addons.current.message.action", "https://www.googleapis.com/auth/gmail.addons.current.message.readonly",
			"https://www.googleapis.com/auth/gmail.addons.current.message.metadata",
			"https://www.googleapis.com/auth/gmail.readonly", "https://www.googleapis.com/auth/gmail.modify"},
		Endpoint: google.Endpoint,
	}
	oauthStateStringGl = ""
)

/*
InitializeOAuthGoogle Function
*/
func InitializeOAuthGoogle() {
	oauthConfGl.ClientID = viper.GetString("google.clientID")
	oauthConfGl.ClientSecret = viper.GetString("google.clientSecret")
	oauthStateStringGl = viper.GetString("oauthStateString")
}

/*
HandleGoogleLogin Function
*/
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	handler.HandleLogin(w, r, oauthConfGl, oauthStateStringGl)
}

/*
CallBackFromGoogle Function
*/
func CallBackFromGoogle(w http.ResponseWriter, r *http.Request) {

	state := r.FormValue("state")
	fmt.Println("state is: ", state)
	if state != oauthStateStringGl {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	fmt.Println("code is: ", code)

	if code == "" {
		w.Write([]byte("Code Not Found to provide AccessToken..\n"))
		reason := r.FormValue("error_reason")
		if reason == "user_denied" {
			w.Write([]byte("User has denied Permission.."))
		}
		// User has denied access..
		// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		token, err := oauthConfGl.Exchange(oauth2.NoContext, code)
		fmt.Println(token.Expiry)
		fmt.Println(token.TokenType)
		fmt.Println()
		fmt.Println("refresh token is --->", token.RefreshToken)
		fmt.Println(token.AccessToken)
		//fmt.Println(token.TokenType)
		//fmt.Println(token.RefreshToken)
		//fmt.Println(token.Expiry)
		//fmt.Println("Access Token: ", token.AccessToken)
		//fmt.Println("expiry: ", token.Expiry)
		//fmt.Println("refrsh token: ", token.RefreshToken)
		//fmt.Println("valid: ", token.Valid())
		//fmt.Println(token.Extra("id_token"))
		if err != nil {
			return
		}

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		defer resp.Body.Close()

		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		//// calender use token instead of tok
		//		tok := &oauth2.Token{AccessToken: token.AccessToken}
		//		client := oauthConfGl.Client(context.Background(), tok)
		//		srv, err := calendar.New(client)
		//		if err != nil {
		//			log.Fatalf("Unable to retrieve Calendar client: %v", err)
		//		}
		//
		//		t := time.Now().Format(time.RFC3339)
		//		events, err := srv.Events.List("primary").ShowDeleted(false).
		//			SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
		//		if err != nil {
		//			log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
		//		}
		//		fmt.Println("Upcoming events:")
		//		if len(events.Items) == 0 {
		//			fmt.Println("No upcoming events found.")
		//		} else {
		//			for _, item := range events.Items {
		//				date := item.Start.DateTime
		//				if date == "" {
		//					date = item.Start.Date
		//				}
		//				fmt.Printf("%v (%v)\n", item.Summary, date)
		//			}
		//		}
		//	// https://content.googleapis.com/calendar/v3/calendars/primary/events?sendNotifications=true&conferenceDataVersion=1&maxAttendees=2&sendUpdates=all&supportsAttachments=false&alt=json&key=AIzaSyAa8yy0GdcGPHdtD083HiGGx_S0vMPScDM
		//	// https://content.googleapis.com/calendar/v3/calendars/primary/events?sendNotifications=true&conferenceDataVersion=1&maxAttendees=2&sendUpdates=all&supportsAttachments=false&alt=json&key=AIzaSyAa8yy0GdcGPHdtD083HiGGx_S0vMPScDM
		//	event := &calendar.Event{
		//			Summary: "Google I/O 2015",
		//			Location: "800 Howard St., San Francisco, CA 94103",
		//			Description: "A chance to hear more about Google's developer products.",
		//			Start: &calendar.EventDateTime{
		//				DateTime: "2021-04-28T09:00:00-07:00",
		//				TimeZone: "Asia/Kolkata",
		//			},
		//			End: &calendar.EventDateTime{
		//				DateTime: "2021-04-28T17:00:00-08:00",
		//				TimeZone: "Asia/Kolkata",
		//			},
		//
		//			//Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
		//			Attendees: []*calendar.EventAttendee{
		//				&calendar.EventAttendee{Email:"ruthala.shiva512@gmail.com", Organizer: false},
		//				&calendar.EventAttendee{Email:"ruthala.charan512@gmail.com", Organizer: false},
		//			},
		//			//Reminders: &calendar.EventReminders{
		//			//	Overrides: []*calendar.EventReminder{
		//			//		{Method: "popup", Minutes: 10},
		//			//	},
		//			//},
		//
		//		}
		//		calendarId := "primary"
		//		event, err = srv.Events.Insert(calendarId, event).SendNotifications(true).SendUpdates("all").Do()
		//		if err != nil {
		//			log.Fatalf("Unable to create event. %v\n", err)
		//		}
		//		fmt.Printf("Event created: %s\n", event.HtmlLink)

		t, err := template.New("webpage").Parse(handler.TokenPage)
		if err != nil {
			log.Fatal(err)
		}

		// data := map[string]interface{}{
		// 	"Token": token.RefreshToken,
		// }

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// w.Write([]byte("Hello, I'm protected\n"))
		// w.Write([]byte(string(response)))
		fmt.Println(t.Execute(w, token))

		return
	}
}
