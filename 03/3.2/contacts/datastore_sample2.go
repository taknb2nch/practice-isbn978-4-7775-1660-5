package contacts

import (
	"fmt"
	"net/http"
	"appengine"
	"appengine/datastore"
)

type Person struct {
	Gender string
	Tel string
}

func init() {
	http.HandleFunc("/writer", writer)
}

func writer(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	//
	key := datastore.NewKey(c, "Contacts", "Tanaka", 0, nil)
	//key := datastore.NewKey(c, "Contacts", "", 1, nil)

	entity := Person {
		Gender: "female",
		Tel: "090-****-xxxx"}

	wkey, err := datastore.Put(c, key, &entity)

	if err != nil {
		c.Infof("%s", err.Error())
		return
	}

	fmt.Fprintf(w, "Stored completed. key: %s", wkey.String())
}