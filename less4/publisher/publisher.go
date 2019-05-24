package main

import (
	"net/http"
	"log"
	"github.com/sanjid133/opentracing-go/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func main()  {
	tracer, closer := tracing.Init("publisher")
	defer closer.Close()

	http.HandleFunc("/publish", func(writer http.ResponseWriter, request *http.Request) {
		spanCtx, _ := tracer.Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(request.Header),
		)
		span := tracer.StartSpan("publish",
			ext.RPCServerOption(spanCtx))
		defer span.Finish()

		str := request.FormValue("str")
		println(str)
		//span.LogKV("value", str)
	})

	println("Running Publisher at :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
