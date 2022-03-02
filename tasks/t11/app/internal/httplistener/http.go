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
func HttpListener(ctx context.Context, http_to_db chan based.Data_to_db, db_to_http chan []byte) {
	log.Println("\t[HTTP] HANDLER STARTED")
	var b based.Data_to_db
	server := http.Server{}
	server.Addr = fmt.Sprintf("%s:%s", cfg.HTTP_HOST, cfg.HTTP_PORT)

	log.Printf("[HTTP] launching server on %s:%s", cfg.HTTP_HOST, cfg.HTTP_PORT)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.Method)
		log.Println(r.URL)
		r.ParseForm()
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "POST" {
			user := r.Form.Get("user_id")
			date := r.Form.Get("date")
			name := r.Form.Get("name")
			if user == "" || date == "" || name == "" {
				log.Println("WRONG REQUEST, no user/date/name", user, date, name)
				fmt.Fprintf(w, "%s", `{"error" : 400}`)
				return
			}
			b.User_id = user
			b.Date = date
			b.Name = name
			if r.URL.String() == "/create_event" {
				log.Println("creating event", r.Form)
				b.Mode = "create"
				http_to_db <- b
				res := <-db_to_http
				fmt.Fprintf(w, "%s", res)
			} else if r.URL.String() == "/update_event" {
				set_user := r.Form.Get("set_user")
				set_date := r.Form.Get("set_date")
				set_name := r.Form.Get("set_name")
				if set_date == "" || set_name == "" || set_user == "" {
					log.Println("wrong request in update, no user/name/date", set_user, set_name, set_date)
					fmt.Fprintf(w, "%s", `{"error" : 400}`)
					return
				}
				log.Println("updating event", r.Form)
				b.Set_date = set_date
				b.Set_name = set_name
				b.Set_user = set_user
				b.Mode = "update"
				http_to_db <- b
				res := <-db_to_http
				fmt.Fprintf(w, "%s", res)
			} else if r.URL.String() == "/delete_event" {
				log.Println("deleting event", r.Form)
				b.Mode = "delete"
				http_to_db <- b
				res := <-db_to_http
				fmt.Fprintf(w, "%s", res)
			} else {
				log.Println("Wrong adress of POST request, something else")
				fmt.Fprintf(w, "%s", `{"error" : 500}`)
			}
		} else if r.Method == "GET" {
			//log.Println(r.URL.Fragment)
			values := r.URL.Query()
			user := values.Get("user_id")
			date := values.Get("date")
			//name := values.Get("name")
			//b.Name = name
			b.User_id = user
			b.Date = date
			if user == "" || date == "" {
				log.Println("WRONG REQUEST of Get, no user/name/date", user, date)
				fmt.Fprintf(w, "%s", `{"error" : 400}`)
				return
			}
			if r.URL.Path == "/events_for_day" {
				b.Mode = "day"
				http_to_db <- b
				res := <-db_to_http
				fmt.Fprintf(w, "%s", res)
			} else if r.URL.Path == "/events_for_week" {
				b.Mode = "week"
				http_to_db <- b
				res := <-db_to_http
				fmt.Fprintf(w, "%s", res)
			} else if r.URL.Path == "/events_for_month" {
				b.Mode = "month"
				http_to_db <- b
				res := <-db_to_http
				fmt.Fprintf(w, "%s", res)
			} else {
				log.Println("wrong Get Method")
				fmt.Fprintf(w, "%s", `{"error" : 500}`)
			}
		} else {
			log.Println("wrong method")
		}
	})

	go func() {
		err := server.ListenAndServe()
		based.CheckErr(err, "error while listening server")
	}()
	<-ctx.Done()
	err := server.Shutdown(context.Background())
	based.CheckErr(err, "error while finishing server")
	log.Println("[HTTP]\tHANDLER FINISHED")
}
