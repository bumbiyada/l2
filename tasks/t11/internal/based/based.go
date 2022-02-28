package based

import (
	"fmt"
	"log"
)

type data_to_db struct {
	user_id string
	date    string
}

func Based() {
	fmt.Println("IT`S BASED")
}

func CheckErr(e error, description string) {

	if e != nil {
		log.Fatalf("\t[ERROR]: \t%s\n\t\t\t[DESCRIPTION]: \t%s", e.Error(), description)
	}
}

// l0

func Get_data(user, data, mode string) (response []byte) {
	switch mode {
	case "day":
		log.Println("get data for day")
	case "week":
		log.Println("get data for week")
	case "month":
		log.Println("get data for month")
	}
	return []byte("data")
}

func Post_data(user, data, mode string) (response []byte) {
	switch mode {
	case "create":
		log.Println("create")
	case "update":
		log.Println("update")
	case "delete":
		log.Println("delete")
	}
	return []byte("data")
}
