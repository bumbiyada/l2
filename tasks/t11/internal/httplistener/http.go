package httplistener

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bumbiyada/l2/tasks/t11/internal/based"
	"github.com/bumbiyada/l2/tasks/t11/internal/cfg"
)

//based.CheckErr()
// function that listens http requests and make control
func HttpListener(ctx context.Context, http_to_db chan based.data_to_db, db_to_http chan []byte) {
	based.Based()
	log.Printf("[HTTP] HOST %s PORT %s", cfg.HTTP_HOST, cfg.HTTP_PORT)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		log.Println(r.URL)
		r.ParseForm()

		if r.Method == "POST" {
			user := r.Form.Get("user_id")
			date := r.Form.Get("data")
			if user == "" || date == "" {
				log.Println("WRONG REQUEST")
				return
			}
			if r.URL.String() == "/create_event" {
				log.Println("creating event", r.Form)
				//Post_data(r.Form[][])
			} else if r.URL.String() == "/update_event" {
				log.Println("creating event", r.Form)
			} else if r.URL.String() == "/delete_event" {
				log.Println("deleting event", r.Form)
			} else {
				log.Println("Wrong adress of POST request")
			}
		} else if r.Method == "GET" {
			//log.Println(r.URL.Fragment)
			values := r.URL.Query()
			user := values.Get("user_id")
			date := values.Get("date")
			if user == "" || date == "" {
				log.Println("WRONG REQUEST")
				return
			}
			if r.URL.Path == "/events_for_day" {

			} else if r.URL.Path == "/events_for_week" {

			} else if r.URL.Path == "/events_for_month" {

			} else {
				log.Println("wrong Get Method")
			}
		} else {
			log.Println("wrong method")
		}
	})

	// if r.Method == http.MethodPost {
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	r.ParseForm()
	// 	log.Printf("[HTTP] GET REQUEST FROM CLIENT %s", r.Form)
	// 	//answer
	// 	fmt.Fprintf(w, "%s", "response")

	// } else {
	// 	log.Println("NOT POST REQUEST")
	// 	fmt.Fprintf(w, "%s", "WHO ARE YOU ?")
	// }
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.HTTP_HOST, cfg.HTTP_PORT), nil))

}
