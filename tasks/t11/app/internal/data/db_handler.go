package data

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/bumbiyada/l2/tasks/t11/internal/based"
	"github.com/bumbiyada/l2/tasks/t11/internal/cfg"
	sqlx "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS event  (
	event_id SERIAL PRIMARY KEY,
	user_id int4,
	name VARCHAR(32),
	event_date DATE
  );
`

type event struct {
	User_id    int64     `json:"user_id" db:"user_id"`
	Name       string    `json:"name" db:"name"`
	Event_date time.Time `json:"date" db:"event_date"`
}

type good_answer struct {
	Result []event `json:"result"`
}

// complete url
var url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DB_USER, cfg.DB_PASS, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)

func Init_db() (*sqlx.DB, error) {

	Db, err := sqlx.Connect("postgres", url)
	based.CheckErr(err, "error while connecting to Database")
	// init db
	Db.MustExec(schema)
	return Db, err
}
func DB_handler(ctx context.Context, http_to_db chan based.Data_to_db, db_to_http chan []byte) {
	log.Println("[DB] HANDLER STARTED")
	if cfg.DEBUG {

		for {
			select {
			case <-ctx.Done():
				log.Println("DB HANDLER CLOSING")
				return
			case o := <-http_to_db:
				log.Println("db got", o)
				db_to_http <- []byte("ok, status 200")
			default:
				time.Sleep(time.Millisecond * 200)
			}
		}
	} else {
		db, err := Init_db()
		based.CheckErr(err, "error")
		for {
			select {
			case <-ctx.Done():
				log.Println("done")
			case o := <-http_to_db:

				switch o.Mode {
				case "create":
					db.MustExec("INSERT INTO event (user_id, name, event_date) VALUES ($1, $2, $3)",
						o.User_id, o.Name, o.Date)
					db_to_http <- []byte(`{"result":200}`)
					log.Println("created new event")
				case "update":
					db.MustExec("UPDATE event SET user_id = $1, name = $2, event_date = $3 WHERE (user_id = $4 AND name = $5 AND event_date = $6)",
						o.Set_user, o.Set_name, o.Set_date, o.User_id, o.Name, o.Date)
					db_to_http <- []byte(`{"result":200}`)
				case "delete":
					db.MustExec("DELETE FROM event  WHERE (user_id = $1 AND name = $2 AND event_date = $3)",
						o.User_id, o.Name, o.Date)
					db_to_http <- []byte(`{"result":200}`)
				case "day", "week", "month":
					events := []event{}
					var cnst int
					if o.Mode == "day" {
						cnst = 1
					} else if o.Mode == "week" {
						cnst = 7
					} else if o.Mode == "month" {
						cnst = 30
					} else {
						log.Println("wrong")
					}
					Date_2, err := time.Parse("2006-01-02", o.Date)
					if err != nil {
						log.Println("error while parsing date")
						db_to_http <- []byte(`{"error": 400}`)
						break
					}
					Date_2 = Date_2.Add(time.Hour * 24 * time.Duration(cnst))
					based.CheckErr(err, "error while parsing date")
					log.Println(o.Date, Date_2.String()[0:10])
					err = db.Select(&events, "SELECT user_id, name, event_date FROM event WHERE user_id = $1 AND event_date BETWEEN $2 AND $3", o.User_id, o.Date, Date_2.String()[0:10])
					based.CheckErr(err, "error while selectind data from db")
					if len(events) == 0 {
						db_to_http <- []byte(`{"error": 404}`)
						break
					}
					ga := good_answer{events}
					res, err := json.Marshal(ga)
					based.CheckErr(err, "error while encoding json")
					db_to_http <- res

				}
			default:
				//log.Println("do smth")
				time.Sleep(time.Millisecond * 230)
			}
		}
	}

}
