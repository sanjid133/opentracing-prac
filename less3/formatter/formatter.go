package main

import (
	"net/http"
	"fmt"
	"log"
)

func main()  {
	http.HandleFunc("/format", func(writer http.ResponseWriter, request *http.Request) {
		name := request.FormValue("name")
		str := fmt.Sprintf("Hello %s", name)
		writer.Write([]byte(str))
	})

	println("Running Formatter at :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
