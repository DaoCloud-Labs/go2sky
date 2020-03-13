# GO2Sky

[![Build](https://github.com/SkyAPM/go2sky/workflows/Build/badge.svg?branch=master)](https://github.com/SkyAPM/go2sky/actions?query=branch%3Amaster+event%3Apush+workflow%3ABuild)
[![GoDoc](https://godoc.org/github.com/SkyAPM/go2sky?status.svg)](https://godoc.org/github.com/SkyAPM/go2sky)


**GO2Sky** is an instrument SDK library, written in Go, by following [Apache SkyWalking](https://github.com/apache/incubator-skywalking) tracing and metrics formats.

**NOTE**: This repo provided by DaoCloud Labs, and used for DaoCloud DMP only.

protocol refs  `reporter/grpc` dir and `https://gitlab.daocloud.io/microservice/dmp/skywalking/skywalking-data-collect-protocol.git`.

# Installation
```
$ go get -u github.com/DaoCloud-Labs/go2sky
```

The API of this project is still evolving. The use of vendoring tool is recommended.

# Quickstart

By completing this quickstart, you will learn how to trace local methods. For more details, please view 
[the example](example_trace_test.go)

## Configuration

GO2Sky can export traces to Apache SkyWalking OAP server or local logger. In the following example, we configure GO2Sky to export to OAP server, 
which is listening on `oap-skywalking` port `11800`, and all of the spans from this program will be associated with a service name `example`.
 
 ```go
r, err := reporter.NewGRPCReporter("oap-skywalking:11800")
if err != nil {
    log.Fatalf("new reporter error %v \n", err)
}
defer r.Close()
tracer, err := go2sky.NewTracer("example", go2sky.WithReporter(r))
```

## Create span

To create a span in a trace, we used the `Tracer` to start a new span. We indicate this as the root span because of 
passing `context.Background()`. We must also be sure to end this span, which will be show in [End span](#end-span).

```go
span, ctx, err := tracer.CreateLocalSpan(context.Background())
```

## Get Global TraceID

Get the `TraceID` of the `activeSpan` in the `Context`.

```go
go2sky.TraceID(ctx)
```

## Create a sub span

A sub span created as the children of root span links to its parent with `Context`.

```go
subSpan, newCtx, err := tracer.CreateLocalSpan(ctx)
```

## End span

We must end the spans so they becomes available for sending to the backend by a reporter.

```go
subSpan.End()
span.End()
```

# Advanced Concepts

We cover some advanced topics about GO2Sky.

## Context propagation

Trace links spans belong to it by using context propagation which varies between different scenario.

### In process

We use `context` package to link spans. The root span usually pick `context.Background()`, and sub spans
will inject the context generated by its parent.

```go
//Create a new context
entrySpan, entryCtx, err := h.tracer.CreateEntrySpan(context.Background(), ...)

// Some operation
...

// Link two spans by injecting entrySpan context into exitSpan
exitSpan, err := t.tracer.CreateExitSpan(entryCtx, ...)

```

### Crossing process

We use `Entry` span to extract context from downstream service, and use `Exit` span to inject context to
upstream service.

`Entry` and `Exit` spans make sense to OAP analysis which generates topology map and service metrics.

```go
//Extract context from HTTP request header `sw6`
span, ctx, err := h.tracer.CreateEntrySpan(r.Context(), "/api/login", func() (string, error) {
		return r.Header.Get("sw6"), nil
})

// Some operation
...

// Inject context into HTTP request header `sw6`
span, err := t.tracer.CreateExitSpan(req.Context(), "/service/validate", "tomcat-service:8080", func(header string) error {
		req.Header.Set(propagation.Header, header)
		return nil
})

```

## Tag

We set tags into a span which is stored in the backend, but some tags have special purpose. OAP server
may use them to aggregate metrics, generate topology map and etc.

They are defined as constant in root package with prefix `Tag`.

## Plugins

Plugins is integrated with specific framework, for instance, `net/http`, `gin` and etc. They
are stored in `plugins` package.

 1. [HTTP client/server example](plugins/http/example_http_test.go)
 1. [gin example](plugins/gin/example_gin_test.go)

# License
Apache License 2.0. See [LICENSE](LICENSE) file for details.
