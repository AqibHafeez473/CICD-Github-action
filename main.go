200~package main

import (
		"fmt"
			"net/http"
		)

		func handler(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello Aqib! Go CI/CD pipeline is working 🚀")
			}

			func main() {
					http.HandleFunc("/", handler)
						http.ListenAndServe(":8080", nil)
					}

