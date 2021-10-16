package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Metric struct {
	Name    string
	Value   string
	Help    string
	Type    string
	Comment string
	Labels  map[string]string
}

func main() {

	PORT := 8090
	requestCounter := 0
	mux := http.NewServeMux()
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusInternalServerError) })
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
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		requestCounter++
		metric := Metric{
			Name:    "request_counter",
			Value:   strconv.Itoa(requestCounter),
			Help:    "request_counter empty",
			Type:    "counter",
			Comment: "Number of requests received",
			Labels: map[string]string{
				"": "",
			},
		}
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte(fmt.Sprintf(`%s`, metricToString(&metric))))
	})

	fmt.Printf("listening on 0.0.0.0:%d ...\n", PORT)
	log.Fatal(http.ListenAndServe("0.0.0.0:8090", mux))
}

func metricToString(m *Metric) string {
	return fmt.Sprintf(
		`# HELP %s
# TYPE %s %s
# Comment %s
%s %s`, m.Help, m.Name, m.Type, m.Comment, m.Name, m.Value)
}
