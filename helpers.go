package main

import "log"

func quitAppIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
