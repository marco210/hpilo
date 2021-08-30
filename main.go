package main

import (
	"fmt"
	"hpilo_exporter/config"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stmcginnis/gofish"
)

func metrichandler(w http.ResponseWriter, r *http.Request) {
	//response, err := http.Get("http://localhost:8080/static/test.json")
	// var err error
	conf := gofish.ClientConfig{
		Endpoint: r.URL.Query().Get("https://192.169.2.2"),
		// Endpoint: r.URL.Query().Get("localhost:8080/static/test.json"),
		Username: "username",
		Password: "password",
		Insecure: true,
	}

	fmt.Println(r.URL.Query().Get("https://192.169.2.2"))

	var err error
	config.GOFISH, err = gofish.Connect(conf)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println("Connected")
	}
	defer config.GOFISH.Logout()

	fmt.Println(" Connect successful")

	mhandler := promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			ErrorHandling: promhttp.ContinueOnError,
		})
	mhandler.ServeHTTP(w, r)
}

func main() {
	const PORT = "9000"
	fmt.Println("Server listening at ", PORT)

	// Listen all interfaces at port 9000
	const IP_ADDRESS = ":" + PORT

	// system := collector.SystemCollector{}
	// fmt.Printf("%v", system)
	// prometheus.Register(system)

	// chassis := chassis.Chassis{}
	// prometheus.Register(chassis)

	// Starting server
	http.HandleFunc("/metrics", metrichandler)
	http.ListenAndServe(IP_ADDRESS, nil)
}
