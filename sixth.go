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
	 v1 := r.Header.Get("HEADER_KEY")
	 headerstr:=string(v1)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
    }

    body := string(reqBody)
    finalres:="BODY:"+body+"HEADER:"+headerstr+"sixth funciton\n"

    w.Write([]byte(finalres))
}
