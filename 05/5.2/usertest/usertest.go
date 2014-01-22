package usertest

import (
	"fmt"
	"net/http"
	"text/template"
	"appengine"
	"appengine/user"
)

type TEMPSUB struct {
	Email string
	Url string
}

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/logined", logined)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	login_url, err := user.LoginURL(c, "/logined")

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t_data := &TEMPSUB {
		Url: login_url,
	}
	tmpl := template.Must(template.New("ROOT").Parse(root_template))
	tmpl.Execute(w, t_data)
}

var root_template = `
<!doctype html>
<html lang="ja">
<body>
	<a href="{{.Url}}">login</a>
</body>
</html>
`
func logined(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	u := user.Current(c)

	logout_url, err := user.LogoutURL(c, "/")

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	tmpl_data := &TEMPSUB {
		Email: u.Email,
		Url: logout_url,
	}

	tmpl := template.Must(template.New("Logined").Parse(logined_template))
	tmpl.Execute(w, tmpl_data)
}

var logined_template = `
<!doctype html>
<html lang="ja">
<body>
  <p>your email: {{.Email}}</p>
  <a href="{{.Url}}">logout</a>
</body>
</html>
`
