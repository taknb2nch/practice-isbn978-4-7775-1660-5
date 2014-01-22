package usertest2

import (
	"fmt"
	"net/http"
	"text/template"
	"appengine"
	"appengine/user"
)

type TEMPSUB struct {
	User string
	Url string
}

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logined", logined)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, root_template)
}

var root_template = `
<!doctype html>
<html lang="ja">
<body>
	<form action="/login"  method="post">
		<input type="text" name="identity" value="input identity" />
		<input type="submit" value="post" />
	</form>
</body>
</html>
`

func login(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	identity := r.FormValue("identity")

	login_url, err := user.LoginURLFederated(c, "/logined", identity)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	http.Redirect(w, r, login_url, http.StatusFound)
}

func logined(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	u := user.Current(c)

	if u == nil {
		// not logined
		http.Redirect(w, r, "/", http.StatusFound)
	}

	logout_url, err := user.LogoutURL(c, "/")

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	tmpl_data := &TEMPSUB {
		User: u.FederatedIdentity,
		Url: logout_url,
	}

	tmpl := template.Must(template.New("Logined").Parse(logined_template))
	tmpl.Execute(w, tmpl_data)
}

var logined_template = `
<!doctype html>
<html lang="ja">
<body>
  <p>your identity: {{.User}}</p>
  <a href="{{.Url}}">logout</a>
</body>
</html>
`
