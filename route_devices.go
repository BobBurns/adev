package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

//query amazon and display device statistics
func devHandler(querys []MetricQuery) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := webSession(w, r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
		for i, _ := range querys {

			err := querys[i].getStatistics("-10m")
			if err != nil {
				log.Printf("Error with getStatistics: %s", err)
				http.Redirect(w, r, "/html/error.html", http.StatusFound)

			}
		}
		var b bytes.Buffer
		err = t.ExecuteTemplate(&b, "home2.html", querys)
		if err != nil {
			fmt.Fprintf(w, "Error with template: %s ", err)
			return
		}
		b.WriteTo(w)

	}
}
