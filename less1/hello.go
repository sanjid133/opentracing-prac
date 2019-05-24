package main

import (
	"os"
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"github.com/sanjid133/opentracing-go/tracing"
)

func main()  {
	if len(os.Args) != 2 {
		panic("Expecting one arg")
	}
	tracer, closer := tracing.Init("hello-world")
	defer closer.Close()

	name := os.Args[1]
	str := fmt.Sprintf("Hello %s", name)

	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", name)

	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", name),
	)
	println(str)
	span.LogKV("event", "println")
	span.Finish()
}



