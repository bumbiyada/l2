package data

import (
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
	description VARCHAR(255),
	date DATE
  );
`

// complete url
var url = "url" + cfg.DB_HOST

func init_db() (*sqlx.DB, error) {

	Db, err := sqlx.Connect("postgres", url)
	based.CheckErr(err, "error while connecting to Database")
	// init db
	Db.MustExec(schema)
	return Db, err
}
