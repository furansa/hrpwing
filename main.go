package main

import (
	"flag"
	"log"
	"net/http"

	"./hrpwing"
)

func main() {
	var debugMode = flag.Bool("debug", false, "Outputs debug information")
	var httpsMode = flag.Bool("https", false, "Uses HTTPS instead of default HTTP")
	flag.Parse()

	p := &hrpwing.Proxy{
		Client:  http.DefaultClient,
		BaseURL: "https://www.golang.org",
	}

	http.Handle("/", p)

	if *httpsMode {
		if *debugMode {
			log.Printf("Listening on port :8443, all request will be redirected to %s", p.BaseURL)
		}

		err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)

		panic(err)
	} else {
		if *debugMode {
			log.Printf("Listening on port :8000, all request will be redirected to %s", p.BaseURL)
		}

		err := http.ListenAndServe(":8000", nil)

		panic(err)
	}
}
