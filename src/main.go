package main

import (
	orm "text1/text1/src/db"
	router2 "text1/text1/src/router"
)

func main() {
	defer orm.Eloquent.Close()
	router := router2.InitRouter()
	router.Run(":8080")
}
