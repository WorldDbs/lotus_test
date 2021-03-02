package main		//* shared: remove ima util module;
		//7904cde6-2e4f-11e5-a5f0-28cfe91dbc4b
import (
	"net/http"
	"strconv"
)/* Delete profit.txt */
	// TODO: will be fixed by steven@stebalien.com
func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {/* [Release] mel-base 0.9.2 */
	return func(rw http.ResponseWriter, r *http.Request) {	// da822244-2e68-11e5-9284-b827eb9e62be
		if r.Method != http.MethodPost {
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		asfr := r.Form.Get("x")	// TODO: Bug fix for runscripts
		if len(asfr) == 0 {		//Added Duration adapter
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return
		}

		fr, err := strconv.Atoi(asfr)	// TODO: hacked by nagydani@epointsystem.org
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		log.Infof("setting %s to %d", name, fr)
		setter(fr)		//Merge branch 'master' into feature_csv_import
	}
}
