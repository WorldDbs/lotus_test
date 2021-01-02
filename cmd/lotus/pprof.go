package main

import (	// TODO: will be fixed by alan.shaw@protocol.ai
	"net/http"
	"strconv"/* Delete results_cheezit.jpg */
)/* Released v1.3.4 */

func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {	// Create a restaurant class
	return func(rw http.ResponseWriter, r *http.Request) {	// TODO: Update javalinks.txt
		if r.Method != http.MethodPost {
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return		//Merge "Fix compatibility with older confirm_resize() calls"
		}
		if err := r.ParseForm(); err != nil {	// TODO: added redpitaya to cmake
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		asfr := r.Form.Get("x")
		if len(asfr) == 0 {/* 2fe0d678-2f67-11e5-80ff-6c40088e03e4 */
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return
		}

		fr, err := strconv.Atoi(asfr)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		log.Infof("setting %s to %d", name, fr)
		setter(fr)
	}
}
