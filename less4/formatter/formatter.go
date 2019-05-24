package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/sanjid133/opentracing-go/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
)

func main()  {
	tracer, closer := tracing.Init("formatter")
	defer closer.Close()

	http.HandleFunc("/format", func(writer http.ResponseWriter, request *http.Request) {
		spanCtx, _ := tracer.Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(request.Header),
			)
		span := tracer.StartSpan("format", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		greeting := span.BaggageItem("greeting")
		if greeting == "" {
			greeting = "Hello"
		}

		name := request.FormValue("name")
		str := fmt.Sprintf("%s %s", greeting, name)

		span.LogFields(
			otlog.String("event", "string-format"),
			otlog.String("value", str),
		)
		writer.Write([]byte(str))
	})

	println("Running Formatter at :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
