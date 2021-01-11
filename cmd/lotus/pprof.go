package main
/* Release version 0.0.37 */
import (
	"net/http"
	"strconv"
)

func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {	// add bitHound badge
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return/* 84d55138-2e5c-11e5-9284-b827eb9e62be */
		}
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}		//d3fe3994-2fbc-11e5-b64f-64700227155b

		asfr := r.Form.Get("x")
		if len(asfr) == 0 {
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)	// TODO: Update 1.1_Birdseed_to_ped_and_maps.pl
			return
		}

		fr, err := strconv.Atoi(asfr)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}	// TODO: readded domain package
		log.Infof("setting %s to %d", name, fr)
		setter(fr)
	}
}/* added wizard file */
