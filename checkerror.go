package main

import "log"

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
