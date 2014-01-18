package bbs

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type Message struct {
	Name    string
	Content string
	Date    time.Time
}

func init() {
	http.HandleFunc("/", view)
	http.HandleFunc("/post", post)
}

func view(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, inputForm)
}

// func post(rw http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(rw, "name1: %s, comment1: %s", req.FormValue("uname"), req.FormValue("comment"))
// }

const inputForm = `
<!doctype html>
<html lang="ja">
  <head>
    <meta charset="utf-8" />
    <title>the bbs</title>
  </head>
  <body>
    <form action="/post" method="post">
      <div>
      	name: <input type="text" name="uname" value="" placeholder="">
      </div>
      <div>
      	comment: <input type="text" name="comment" value="" placeholder="">
      </div>
      <div>
      	<input type="submit" name="submit_name" value="post">
      </div>
    </form>
  </body>
</html>
`

func post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	data := Message{
		Name:    r.FormValue("uname"),
		Content: r.FormValue("comment"),
		Date:    time.Now()}

	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Message", nil), &data)

	if err != nil {
		fmt.Fprint(w, "error")
		return
	}

	q := datastore.NewQuery("Message").Order("-Date").Limit(10)

	messages := make([]Message, 0)

	_, err = q.GetAll(c, &messages)

	if err != nil {
		fmt.Fprint(w, "error")
	}

	//
	c.Infof("[receive] author: %s, content: %s", data.Name, data.Content)

	for _, j := range messages {
		fmt.Fprintf(w, "author: %s \n%s \n\n", j.Name, j.Content)
	}

	//
	tmpl := template.Must(template.New("bbs").Parse(tmplHTML))
	tmpl.Execute(w, messages)
}

const tmplHTML = `
<!doctype html>
<html lang="ja">
  <body>
    {{range .}}
    <div>
      <p>author : <b>{{html .Name}}</b></p>
      <p>{{html .Content}}</p>
    </div>
    {{end}}
  </body>
</html>
`
