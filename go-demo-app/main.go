package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// very basic Metric data object
type Metric struct {
	Name    string
	Value   string
	Help    string
	Type    string
	Comment string
	Labels  map[string]string
}

func main() {
	// port application is listening on
	PORT := 8090

	// in memory request counter
	requestCounter := 0
	mux := http.NewServeMux()

	// explicit route so it doesnt fallback to "/" and raising the counter
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusInternalServerError) })

	// simple landing page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestCounter++
		w.Write([]byte(`
			<html>
				<body>
					<p>Endpoint:</p>
					<ul>
						<li>GET: <a href="/metrics">/metrics</a></li>
					</ul>
				</body>
			</html>	
		`))
	})

	// endpoint for Prometheus to scrape
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		requestCounter++

		// creating Metric
		metric := Metric{
			Name:    "request_counter",             // can be freely choosen
			Value:   strconv.Itoa(requestCounter),  // the value to expose (here the counter)
			Help:    "request_counter empty",       // some additional info that Prometheus collects
			Type:    "counter",                     // on of these: Counter, Gauge, Histogram, Summary, Untyped
			Comment: "Number of requests received", // Prometheus ignores this
			Labels: map[string]string{ // Possible labels, not implemented here
				"": "",
			},
		}
		w.Header().Add("Content-Type", "text/plain")                // set Content-Type Header
		w.Write([]byte(fmt.Sprintf(`%s`, metricToString(&metric)))) // write metric into response after formatting it to a string
	})

	// start server
	fmt.Printf("listening on 0.0.0.0:%d ...\n", PORT)
	log.Fatal(http.ListenAndServe("0.0.0.0:8090", mux))
}

// basic formatting from Metric to string
// for extended functionality use libraries
// https://prometheus.io/docs/guides/go-application/
func metricToString(m *Metric) string {
	return fmt.Sprintf(
		`# HELP %s
# TYPE %s %s
# Comment %s
%s %s`, m.Help, m.Name, m.Type, m.Comment, m.Name, m.Value)
}
