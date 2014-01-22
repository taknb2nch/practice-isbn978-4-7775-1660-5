package crontest

import (
	"fmt"
	"net/http"
	"appengine"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/testjob1", job1)
	http.HandleFunc("/testjob2", job2)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "cron job test")
}

func job1(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("cron job 1 is running.")
}

func job2(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("cron job 2 is running.")
}