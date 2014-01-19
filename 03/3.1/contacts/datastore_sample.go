package contacts

import (
	"fmt"
	"net/http"
	"appengine"
	"appengine/datastore"
)

type Person struct {
	Name string
	Gender string
	Tel string
}

func init() {
	http.HandleFunc("/writer", writer)
}

func writer(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	//
	key := datastore.NewIncompleteKey(c, "Contacts", nil)

	entity := Person {
		Name: "Yoshida",
		Gender: "female",
		Tel: "090-****-xxxx"}

	wkey, err := datastore.Put(c, key, &entity)

	if err != nil {
		c.Infof("%s", err.Error())
		return
	}

	fmt.Fprintf(w, "Stored completed. key: %s", wkey.String())
}