package httptransport

import (
	"net/http"
	"net/url"

	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/tracing/zipkin"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"

	"microservice/pkg/endpoints"
)

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHTTPHandler(endpoints endpoints.Set, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) http.Handler {
	// Zipkin HTTP Server Trace can either be instantiated per endpoint with a
	// provided operation name or a global tracing service can be instantiated
	// without an operation name and fed to each Go kit endpoint as ServerOption.
	// In the latter case, the operation name will be the endpoint's http method.
	// We demonstrate a global tracing service here.
	zipkinServer := zipkin.HTTPServerTrace(zipkinTracer)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		zipkinServer,
	}

	m := http.NewServeMux()
	m.Handle("/sum", httptransport.NewServer(
		endpoints.SumEndpoint,
		decodeHTTPSumRequest,
		encodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "Sum", logger)))...,
	))
	m.Handle("/concat", httptransport.NewServer(
		endpoints.ConcatEndpoint,
		decodeHTTPConcatRequest,
		encodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "Concat", logger)))...,
	))
	return m
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}

type errorWrapper struct {
	Error string `json:"error"`
}
