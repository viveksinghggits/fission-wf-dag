// Due to go plugin mechanism,
// the package of function handler must be main package
package main

import (
	"io/ioutil"
	"net/http"
)

// Handler is the entry point for this fission function
func Handler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	parent := r.Header.Get("Parent")
	body := string(reqBody)
	finalres := body + " " + parent + "\n"
	w.Header().Set("Parent", "Third")
	w.Write([]byte(finalres))
}
