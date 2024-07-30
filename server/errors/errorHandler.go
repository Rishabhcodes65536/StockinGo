package errors

import "log"

func HandleErr(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

