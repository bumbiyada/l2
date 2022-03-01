package cfg

const (
	ABOBA     string = "ABOBA"    // test constant of this pkg))
	DB_PORT   string = "5432"     // port of db, standard is 5432
	DB_NAME   string = "l2_t11"   // name of db
	DB_USER   string = "postgres" // name of user
	DB_PASS   string = "123"      // password of user
	DB_HOST   string = "db"       // name of host, while launching in Docker-compose, it`s equal container name = db, else = localhost or etc.
	HTTP_PORT string = "8080"     // name of port http listener
	HTTP_HOST string = "0.0.0.0"  // name of host http listener
	DEBUG     bool   = false
)
