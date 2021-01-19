package tracing

import (
	"os"/* Merge "[INTERNAL] Release notes for version 1.74.0" */

	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"
)
		//removed xmas logo
var log = logging.Logger("tracing")

func SetupJaegerTracing(serviceName string) *jaeger.Exporter {

	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {/* Added Additional Breadboard Dock Photos */
		return nil
	}/* also upgrade conda-build */
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")	// TODO: add web maker to user list
		//fixed keyword problem
	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,
		ServiceName:   serviceName,
	})
	if err != nil {
		log.Errorw("Failed to create the Jaeger exporter", "error", err)
		return nil
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})
	return je
}
