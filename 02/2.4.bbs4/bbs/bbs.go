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
	Email   string
	Date    time.Time
}

func init() {
	http.HandleFunc("/", view)
	http.HandleFunc("/post", post)
}

func view(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	q := datastore.NewQuery("Message").Order("-Date").Limit(10)

	messages := make([]Message, 0)

	_, err := q.GetAll(c, &messages)

	if err != nil {
		fmt.Fprint(w, "error")
	}

	//
	tmpl := template.Must(template.New("view").Parse(inputTmpl))
	tmpl.Execute(w, messages)
}

const inputTmpl = `
<!doctype html>
<html lang="ja">
  <head>
    <meta charset="utf-8" />
    <title>the bbs</title>
    <link rel="stylesheet" type="text/css" href="css/style.css" />
  </head>
  <body>
    {{range .}}
    <div class="message">
      <p>author : {{html .Name}} Email: {{html .Email}}</p>
      <p>{{html .Content}}</p>
    </div>
    {{end}}
    <div class="post">
    <form action="/post" method="post">
      <div>
      	name: <input type="text" name="uname" value="" placeholder="">
      </div>
      <div>
      	email: <input type="text" name="email" value="" placeholder="">
      </div>
      <div>
      	comment: <input type="text" name="comment" value="" placeholder="">
      </div>
      <div>
      	<input type="submit" name="submit_name" value="post">
      </div>
    </form>
    </div>
  </body>
</html>
`

func post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	data := Message{
		Name:    r.FormValue("uname"),
		Content: r.FormValue("comment"),
		Email:   r.FormValue("email"),
		Date:    time.Now()}

	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Message", nil), &data)

	if err != nil {
		fmt.Fprint(w, "error")
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
