// Due to go plugin mechanism,
// the package of function handler must be main package
package main

import (
    "net/http"
"io/ioutil"
)

// Handler is the entry point for this fission function
func Handler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
    }

    body := string(reqBody)
    finalres:=body+"fifth funciton\n"

    w.Write([]byte(finalres))
}
