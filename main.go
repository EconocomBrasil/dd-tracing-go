package main

import (
	"fmt"
	"log"
	"net/http"
	"net"
	//"github.com/gorilla/mux"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	httptracer "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"time"
)

func main() {
	addr := net.JoinHostPort("localhost","8126")
	tracer.Start(tracer.WithAgentAddr(addr),tracer.WithDebugMode(true))
	defer tracer.Stop()
	r := muxtrace.NewRouter(muxtrace.WithServiceName("go tracing"))

	r.HandleFunc("/gotracing",Metodo).Methods("GET")
	r.HandleFunc("/httpbin",Httpbin).Methods("GET")
	var port = ":3002"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func Metodo(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	span,_ := tracer.StartSpanFromContext(r.Context(), "metodo")
    defer span.Finish()
	fmt.Println("Teste Main Go tracing")
	Httpbin(w,r)
}

func Httpbin(w http.ResponseWriter, r *http.Request){
	client := &http.Client{}
	httpClient := httptracer.WrapClient(client)
	httpClient.Get("https://httpbin.org")
}