package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func HelloParam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if _, ok := params["PARAM"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := w.Write([]byte(fmt.Sprintf("Hello, %s!", params["PARAM"])))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func InternalError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

}
func ParseBody(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(fmt.Sprintf("I got message:\n%s", string(data))))
}
func HeadersSum(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result := a + b
	w.Header().Set("a+b", strconv.Itoa(result))

}
