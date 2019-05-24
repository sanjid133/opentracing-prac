package main

import (
	"net/http"
	"log"
)

func main()  {
	http.HandleFunc("/publish", func(writer http.ResponseWriter, request *http.Request) {
		str := request.FormValue("str")
		println(str)
	})

	println("Running Publisher at :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
