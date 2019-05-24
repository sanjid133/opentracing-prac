package main

import (
	"os"
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"github.com/sanjid133/opentracing-go/tracing"
	"github.com/opentracing/opentracing-go"
	"context"
)

func formatString(ctx context.Context, name string) string {
	span,  _:= opentracing.StartSpanFromContext(ctx, "formatString")
	defer span.Finish()

	str := fmt.Sprintf("Hello %s", name)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", name),
	)
	return str
}

func printHello(ctx context.Context, str string)  {
	span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
	defer span.Finish()

	println(str)
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



