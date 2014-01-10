package bbs

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", view)
	http.HandleFunc("/post", post)
}

func view(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, inputForm)
}

func post(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "name1: %s, comment1: %s", req.FormValue("uname"), req.FormValue("comment"))
}

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