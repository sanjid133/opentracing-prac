package main

import (
	"os"
	"github.com/opentracing/opentracing-go/log"
	"github.com/sanjid133/opentracing-go/tracing"
	"github.com/opentracing/opentracing-go"
	"context"
	"net/url"
	"net/http"
	xhttp "github.com/sanjid133/opentracing-go/util/http"
	"github.com/opentracing/opentracing-go/ext"
)

func formatString(ctx context.Context, name string) string {
	span,  _:= opentracing.StartSpanFromContext(ctx, "formatString")
	defer span.Finish()

	ext.SpanKindRPCClient.Set(span)

	v := url.Values{}
	v.Set("name", name)
	url := "http://localhost:8081/format?"+v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	resp, err := xhttp.Do(req)
	if err != nil {
		panic(err)
	}

	str := string(resp)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", name),
	)
	return str
}

func printHello(ctx context.Context, str string)  {
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer span.Finish()

	ext.SpanKindRPCClient.Set(span)

	v := url.Values{}
	v.Set("str", str)
	url := "http://localhost:8082/publish?"+v.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)
	_, err = xhttp.Do(req)
	if err != nil {
		panic(err)
	}
	span.LogKV("event", "println")
}

func main()  {
	if len(os.Args) != 2 {
		panic("Expecting one arg")
	}
	tracer, closer := tracing.Init("hello-world")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer) //The StartSpanFromContext function uses opentracing.GlobalTracer()

	name := os.Args[1]

	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", name)


	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	hellStr := formatString(ctx, name)
	printHello(ctx, hellStr)
	span.Finish()
}



