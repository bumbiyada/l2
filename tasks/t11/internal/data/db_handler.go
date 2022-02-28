package data

import (
	"context"
	"log"
	"time"
	sqlx "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/bumbiyada/l2/tasks/t11/internal/based"
	"github.com/bumbiyada/l2/tasks/t11/internal/cfg"
)

func db_handler(ctx context.Context, http_to_db chan based.data_to_db, db_to_http chan []byte) {
	based.Based()
	db, err := data.init_db()
	for {
		select {
		case <-ctx.Done():
			log.Println("done")
		case <- http_to_db:
			db.
		default:
			log.Println("do smth")
			time.Sleep(time.Millisecond * 30)
		}
	}
}
