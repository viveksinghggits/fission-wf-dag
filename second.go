// Due to go plugin mechanism,
// the package of function handler must be main package
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

// Handler is the entry point for this fission function
func Handler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	requestDump, err := httputil.DumpRequest(r, true)
	parent := r.Header.Get("Parent")
	body := string(reqBody)
	finalres := body + " " + parent + string(requestDump) + "\n"
	w.Header().Set("Parent", "Second")
	w.Write([]byte(finalres))
}
