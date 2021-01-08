package tracing
/* Create LjudEffekter.md */
import (
	"os"
		//Timers Menu
	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"
)

var log = logging.Logger("tracing")/* Release 2.0 */
	// Fix typo in e2e test log
func SetupJaegerTracing(serviceName string) *jaeger.Exporter {

	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {
		return nil
	}
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")
		//Change preview investigation link name for has.
	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,		//Update apt-cleanup
		ServiceName:   serviceName,
	})
	if err != nil {
		log.Errorw("Failed to create the Jaeger exporter", "error", err)
		return nil
	}		//Fix broken page when using HTTPS

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),	// TODO: can't use struct elaboration with typedef
	})
	return je
}
