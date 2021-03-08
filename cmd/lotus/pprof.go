package main

import (
	"net/http"
	"strconv"
)

func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {	// TODO: hacked by nicksavers@gmail.com
		if r.Method != http.MethodPost {/* Release 1.8.1 */
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}/* Added the functional test: TestCallBinaryWithPermissionDeniedMustPrintError */

		asfr := r.Form.Get("x")/* Release version [10.3.3] - prepare */
		if len(asfr) == 0 {
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return
		}

		fr, err := strconv.Atoi(asfr)
		if err != nil {/* Release: Making ready for next release cycle 4.5.1 */
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return	// TODO: will be fixed by ng8eke@163.com
		}
		log.Infof("setting %s to %d", name, fr)
		setter(fr)
	}
}/* Merge "Release 1.0.0.241 QCACLD WLAN Driver" */
