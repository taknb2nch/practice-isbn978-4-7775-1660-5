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

	entities := make([]Person, 0)

	//
	q := datastore.NewQuery("Contacts").Filter("Gender=", "male")

	keys, err := q.GetAll(c, &entities)

	if err != nil {
		c.Infof("%s", err.Error())
		return
	}

	for i, entity := range entities {
		fmt.Fprintf(w, "%s %s %s\n", keys[i].StringID(), entity.Gender, entity.Tel)
	}
}