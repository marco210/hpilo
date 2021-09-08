package main

import (
	"encoding/json"
	"fmt"
	"hpilo_exporter/collector"
	"hpilo_exporter/config"
	redfishstruct "hpilo_exporter/redfish_struct"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stmcginnis/gofish"
)

var smart_storage redfishstruct.SmartStorage

func metrichandler(w http.ResponseWriter, r *http.Request) {
	// var err error
	conf := gofish.ClientConfig{
		Endpoint: r.URL.Query().Get("ilo_host"),
		Username: "username",
		Password: "password",
		Insecure: true,
	}

	fmt.Println(r.URL.Query().Get("ilo_host"))

	var err error
	config.GOFISH, err = gofish.Connect(conf)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println("Connected")
	}

	defer config.GOFISH.Logout()

	// smartstorage, _ := config.GOFISH.Get("/redfish/v1/Systems/1/SmartStorage/ArrayControllers/0/StorageEnclosures/0")
	smartstorage, _ := config.GOFISH.Get("/redfish/v1/Systems/1/SmartStorage")
	fmt.Println(smartstorage)
	bodyBytes, _ := ioutil.ReadAll(smartstorage.Body)
	//fmt.Println(bodyBytes)
	err = json.Unmarshal(bodyBytes, &smart_storage)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", smart_storage)

	fmt.Printf(smart_storage.Description)

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

	system := collector.SystemCollector{}
	fmt.Printf("%v", system)
	prometheus.Register(system)

	//chassis := collector.Chassis{}
	//prometheus.Register(chassis)
	fmt.Println(IP_ADDRESS)
	// Starting server
	http.HandleFunc("/metrics", metrichandler)
	http.ListenAndServe(IP_ADDRESS, nil)
}
