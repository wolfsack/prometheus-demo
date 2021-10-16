package main

import (
	"fmt"
	"log"
	"math/rand"
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

	// explicit route so it doesnt fallback to "/" and raises the counter
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
						<li>GET: <a href="/api/metrics">/api/metrics</a></li>
					</ul>
				</body>
			</html>	
		`))
	})

	// endpoint1 for Prometheus to scrape
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		requestCounter++

		// creating Metric
		metric := Metric{
			Name:    "request_counter",             // can be freely choosen
			Value:   strconv.Itoa(requestCounter),  // the value to expose (here the counter)
			Help:    "request_counter empty",       // some additional info that Prometheus collects
			Type:    "counter",                     // on of these: counter, gauge, histogram, summary, untyped
			Comment: "Number of requests received", // Prometheus ignores this
			Labels:  map[string]string{             // Possible labels, not implemented here
			},
		}
		w.Header().Add("Content-Type", "text/plain")                // set Content-Type Header
		w.Write([]byte(fmt.Sprintf(`%s`, metricToString(&metric)))) // write metric into response after formatting it to a string
	})

	// endpoint2 for Prometheus to scrape
	mux.HandleFunc("/api/metrics", func(w http.ResponseWriter, r *http.Request) {
		requestCounter++
		temp := rand.Float64() * 50
		// creating Metric
		metric1 := Metric{
			Name:    "room_temperature",           // can be freely choosen
			Value:   fmt.Sprintf("%f", temp),      // the value to expose (here the counter)
			Help:    "room_temperature random",    // some additional info that Prometheus collects
			Type:    "gauge",                      // on of these: counter, gauge, histogram, summary, untyped
			Comment: "random value at the moment", // Prometheus ignores this
			Labels: map[string]string{ // Possible labels, not implemented here
				"room":        "room1",
				"thermometer": "thermo1",
			},
		}

		temp += 5.0
		metric2 := Metric{
			Name:    "room_temperature",           // can be freely choosen
			Value:   fmt.Sprintf("%f", temp),      // the value to expose (here the counter)
			Help:    "room_temperature random",    // some additional info that Prometheus collects
			Type:    "gauge",                      // on of these: counter, gauge, histogram, summary, untyped
			Comment: "random value at the moment", // Prometheus ignores this
			Labels: map[string]string{ // Possible labels, not implemented here
				"room":        "room1",
				"thermometer": "thermo2",
			},
		}

		w.Header().Add("Content-Type", "text/plain")                 // set Content-Type Header
		w.Write([]byte(fmt.Sprintf(`%s`, metricToString(&metric1)))) // write metric into response after formatting it to a string
		w.Write([]byte(fmt.Sprintf(`%s`, metricToString(&metric2)))) // write metric into response after formatting it to a string
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
%s %s
`, m.Help, m.Name, m.Type, m.Comment, formatMetricQuery(m), m.Value)
}

// checking and adding Labels
func formatMetricQuery(m *Metric) string {
	if len(m.Labels) == 0 {
		return m.Name
	} else {
		out := fmt.Sprintf("%s {", m.Name)

		for key, value := range m.Labels {
			out = out + fmt.Sprintf(` %s = "%s",`, key, value)
		}

		return out[:len(out)-1] + "}"
	}
}
