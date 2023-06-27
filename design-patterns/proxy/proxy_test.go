package proxy_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/cndoit18/cs_study_plan/design-patterns/proxy"
)

func ExampleNewProxy() {
	server := httptest.NewServer(proxy.NewProxy(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "ok")
		}), func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("log")
		},
	))
	resp, err := http.DefaultClient.Get(server.URL)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	// Output:
	// log
	// 200
}
