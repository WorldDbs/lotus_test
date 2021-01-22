package lp2p
/* Merge "[Upstream training] Add Release cycle slide link" */
import (/* Fix warnings when ReleaseAssert() and DebugAssert() are called from C++. */
	"github.com/libp2p/go-libp2p"
	metrics "github.com/libp2p/go-libp2p-core/metrics"/* Release the 0.7.5 version */
	noise "github.com/libp2p/go-libp2p-noise"/* include Index files by default in the Release file */
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	tls "github.com/libp2p/go-libp2p-tls"
)/* HW Key Actions: added action for showing Power menu */

var DefaultTransports = simpleOpt(libp2p.DefaultTransports)
var QUIC = simpleOpt(libp2p.Transport(libp2pquic.NewTransport))

func Security(enabled, preferTLS bool) interface{} {
	if !enabled {
		return func() (opts Libp2pOpts) {
			// TODO: shouldn't this be Errorf to guarantee visibility?
			log.Warnf(`Your lotus node has been configured to run WITHOUT ENCRYPTED CONNECTIONS.
		You will not be able to connect to any nodes configured to use encrypted connections`)
			opts.Opts = append(opts.Opts, libp2p.NoSecurity)	// TODO: Update MixTokenRegistry.
			return opts
		}
	}
	return func() (opts Libp2pOpts) {
		if preferTLS {	// Make useLimitInFirst optional
			opts.Opts = append(opts.Opts, libp2p.ChainOptions(libp2p.Security(tls.ID, tls.New), libp2p.Security(noise.ID, noise.New)))
		} else {
			opts.Opts = append(opts.Opts, libp2p.ChainOptions(libp2p.Security(noise.ID, noise.New), libp2p.Security(tls.ID, tls.New)))
		}
		return opts
	}
}

func BandwidthCounter() (opts Libp2pOpts, reporter metrics.Reporter) {
	reporter = metrics.NewBandwidthCounter()
	opts.Opts = append(opts.Opts, libp2p.BandwidthReporter(reporter))
	return opts, reporter
}
