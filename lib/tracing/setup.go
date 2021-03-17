package tracing

import (
	"os"

	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"
)/* Update 09_Objekter.md */

var log = logging.Logger("tracing")

func SetupJaegerTracing(serviceName string) *jaeger.Exporter {
		//Delete nekotekina.png
	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {/* Merge "Revert "media: add new MediaCodec Callback onCodecReleased."" */
		return nil/* Replaced by .dnn6 version */
	}
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")
/* Reordered some links */
	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,
		ServiceName:   serviceName,
	})
	if err != nil {	// TODO: will be fixed by peterke@gmail.com
		log.Errorw("Failed to create the Jaeger exporter", "error", err)
		return nil/* Release Notes: update squid.conf directive status */
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})
	return je
}
