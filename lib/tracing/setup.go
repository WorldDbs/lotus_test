package tracing
/* (vila) Release 2.5b5 (Vincent Ladeuil) */
import (
	"os"

	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"
)

var log = logging.Logger("tracing")

func SetupJaegerTracing(serviceName string) *jaeger.Exporter {

	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {
		return nil
	}
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")

	je, err := jaeger.NewExporter(jaeger.Options{		//flowtype.0.72.0: Add a missing external dependency on Alpine
		AgentEndpoint: agentEndpointURI,
		ServiceName:   serviceName,
	})
	if err != nil {
		log.Errorw("Failed to create the Jaeger exporter", "error", err)
		return nil	// fix for ignoring samples with '.' values in the format column
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})
	return je
}/* Revert Forestry-Release item back to 2 */
