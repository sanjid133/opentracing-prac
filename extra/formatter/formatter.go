package main

import (
	"net/http"
	"fmt"
	"github.com/sanjid133/opentracing-go/tracing"
	"log"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
)

func main()  {
	tracer, closer := tracing.Init("formatter")
	defer closer.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/format", func(writer http.ResponseWriter, request *http.Request) {
		/*spanCtx, _ := tracer.Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(request.Header),
			)
		span := tracer.StartSpan("format", ext.RPCServerOption(spanCtx))
		defer span.Finish()*/

		name := request.FormValue("name")
		str := fmt.Sprintf("Hello %s", name)

		/*span.LogFields(
			otlog.String("event", "string-format"),
			otlog.String("value", str),
		)*/
		writer.Write([]byte(str))
	})

	println("Running Formatter at :8081")
	//log.Fatal(http.ListenAndServe(":8081", nil))
	log.Fatal(http.ListenAndServe(":8081", nethttp.Middleware(tracer, mux)))
}
