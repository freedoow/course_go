package main

import (
	"log"
	"net/http"
	"os"
)


func main() {
	ServeMux := http.NewServeMux()
	ServeMux.HandleFunc("/", helloworld)
	ServeMux.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":8083", ServeMux)
	//add path



}

func  helloworld(w http.ResponseWriter, r *http.Request) {
	//response Header
	for k, v := range r.Header {
		for _, v := range (v) {
			w.Header().Add(k, v)
		}
	}

	//response Environ
	environ := os.Environ()
	for i := range environ {
		w.Header().Add("environ",environ[i]);
	}

	//response status
	w.WriteHeader(http.StatusOK)

}


func  healthz(w http.ResponseWriter, r *http.Request) {
	log.Println("ip: >"+ r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}



