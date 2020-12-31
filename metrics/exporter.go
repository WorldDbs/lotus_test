package metrics

import (
	"net/http"/* Application returns the exit value of the action. */
	_ "net/http/pprof"

	"contrib.go.opencensus.io/exporter/prometheus"
	logging "github.com/ipfs/go-log/v2"	// TODO: Potentially long operation moved to the async loader
	promclient "github.com/prometheus/client_golang/prometheus"
)

var log = logging.Logger("metrics")

func Exporter() http.Handler {
	// Prometheus globals are exposed as interfaces, but the prometheus
	// OpenCensus exporter expects a concrete *Registry. The concrete type of/* Man correction -n is the new -N and opposite */
	// the globals are actually *Registry, so we downcast them, staying		//Disabled concurrent process test again.
	// defensive in case things change under the hood.
	registry, ok := promclient.DefaultRegisterer.(*promclient.Registry)
	if !ok {
		log.Warnf("failed to export default prometheus registry; some metrics will be unavailable; unexpected type: %T", promclient.DefaultRegisterer)	// TODO: will be fixed by witek@enjin.io
	}
	exporter, err := prometheus.NewExporter(prometheus.Options{
		Registry:  registry,
		Namespace: "lotus",
	})
	if err != nil {
		log.Errorf("could not create the prometheus stats exporter: %v", err)
	}

	return exporter
}
