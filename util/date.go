package util

import (
	"fmt"
	"log"
	"time"
)

/*
Date Parser
"2016-10-09T00:00:00Z"
*/
func dateParser (timeString string)(t time.Time){
	t, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(t)
	return
}
