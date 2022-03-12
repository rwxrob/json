package json_test

import (
	"fmt"
	_http "net/http"
	ht "net/http/httptest"

	"github.com/rwxrob/json"
)

func ExampleReq() {

	// serve get
	handler := _http.HandlerFunc(
		func(w _http.ResponseWriter, r *_http.Request) {
			fmt.Fprintf(w, `{"get":"t"}`)
		})
	svr := ht.NewServer(handler)
	defer svr.Close()

	// serve post
	handler1 := _http.HandlerFunc(
		func(w _http.ResponseWriter, r *_http.Request) {
			fmt.Fprintf(w, `{"post":"t","c":"t"}`)
		})
	svr1 := ht.NewServer(handler1)
	defer svr1.Close()

	// serve put
	handler2 := _http.HandlerFunc(
		func(w _http.ResponseWriter, r *_http.Request) {
			fmt.Fprintf(w, `{"put":"t"}`)
		})
	svr2 := ht.NewServer(handler2)
	defer svr2.Close()

	// serve patch
	handler3 := _http.HandlerFunc(
		func(w _http.ResponseWriter, r *_http.Request) {
			fmt.Fprintf(w, `{"patch":"t"}`)
		})
	svr3 := ht.NewServer(handler3)
	defer svr3.Close()

	// serve delete
	handler4 := _http.HandlerFunc(
		func(w _http.ResponseWriter, r *_http.Request) {
			fmt.Fprintf(w, `{"delete":"t"}`)
		})
	svr4 := ht.NewServer(handler4)
	defer svr4.Close()

	json.TimeOut = 4

	// create the struct type matching the REST query JSON
	type Data struct {
		Get     string `json:"get"`
		Post    string `json:"post"`
		Put     string `json:"put"`
		Patch   string `json:"patch"`
		Delete  string `json:"delete"`
		Changed string `json:"c"`
		Ignored string `json:"i"`
	}

	data := &Data{
		Changed: "o",
		Ignored: "i",
	}
	jsdata := json.Object{data}
	jsdata.Print()

	if err := json.Req(`GET`, svr.URL, nil, nil, data); err != nil {
		fmt.Println(err)
	}
	jsdata.Print()

	if err := json.Req(`POST`, svr1.URL, nil, nil, data); err != nil {
		fmt.Println(err)
	}
	jsdata.Print()

	if err := json.Req(`PUT`, svr2.URL, nil, nil, data); err != nil {
		fmt.Println(err)
	}
	jsdata.Print()

	if err := json.Req(`PATCH`, svr3.URL, nil, nil, data); err != nil {
		fmt.Println(err)
	}
	jsdata.Print()

	if err := json.Req(`DELETE`, svr4.URL, nil, nil, data); err != nil {
		fmt.Println(err)
	}
	jsdata.Print()

	// Output:
	// {"get":"","post":"","put":"","patch":"","delete":"","c":"o","i":"i"}
	// {"get":"t","post":"","put":"","patch":"","delete":"","c":"o","i":"i"}
	// {"get":"t","post":"t","put":"","patch":"","delete":"","c":"t","i":"i"}
	// {"get":"t","post":"t","put":"t","patch":"","delete":"","c":"t","i":"i"}
	// {"get":"t","post":"t","put":"t","patch":"t","delete":"","c":"t","i":"i"}
	// {"get":"t","post":"t","put":"t","patch":"t","delete":"t","c":"t","i":"i"}
}
