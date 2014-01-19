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
	http.HandleFunc("/reader", reader)
}

func writer(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	//
	key := datastore.NewKey(c, "Contacts", "Tanaka", 0, nil)

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

func reader(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	entity := new(Person)

	key := datastore.NewKey(c, "Contacts", "Tanaka", 0, nil)

	err := datastore.Get(c, key, entity)

	if err != nil {
		c.Infof("%s", err.Error())
		return
	}

	fmt.Fprintf(w, "%s %s %s\n", key.StringID(), entity.Gender, entity.Tel)
}