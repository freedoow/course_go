package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)


// TODO 作业第三点可写全局拦截器进行日志处理  待学习
func main() {
	ServeMux := http.NewServeMux()
	ServeMux.HandleFunc("/", root)
	ServeMux.HandleFunc("/getAllHeader", getAllHeader)
	ServeMux.HandleFunc("/getEnviron", getEnviron)
	ServeMux.HandleFunc("/recordInfo", recordInfo)
	ServeMux.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":80", ServeMux)
	//add path

}

func  root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	fmt.Println("ip: >"+ r.RemoteAddr)
	// ip
	fmt.Fprint(w,  fmt.Sprintf("ip-->\n%s\n", r.RemoteAddr))
}

func  getAllHeader(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w,  "循环输出heard------>\n\n")
	//
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
		fmt.Fprint(w,  "\n")

	}
}

func  getEnviron(w http.ResponseWriter, r *http.Request) {
	environ := os.Environ()
	for i := range environ {
		fmt.Println(environ[i])
		io.WriteString(w, fmt.Sprintf("PATH is %s\n",environ[i]))
		fmt.Fprint(w,  "\n")
	}
	fmt.Println("**************************")
}

func  recordInfo(w http.ResponseWriter, r *http.Request) {
	// ip
	fmt.Fprint(w,  fmt.Sprintf("ip-->\n%s\n", r.RemoteAddr))
}



func  healthz(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,  "200")
}

