package taskqueuetest

import (
	"fmt"
	"net/http"
	"appengine"
	"appengine/taskqueue"
	"net/url"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/testjob", job)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	fmt.Fprintf(w, "taskqueue test")

	//t := taskqueue.NewPOSTTask("/testjob", nil)
	t := taskqueue.NewPOSTTask("/testjob", url.Values {
		"message": []string{"hello TaskQueue!"},
		})
	taskqueue.Add(c, t, "")
}

func job(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("task queue is running.")
	c.Infof("message -> %s", r.FormValue("message"))
}