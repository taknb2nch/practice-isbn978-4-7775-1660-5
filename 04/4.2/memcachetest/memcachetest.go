package memcachetest

import (
	"fmt"
	"net/http"
	"appengine"
	"appengine/memcache"
)

func init() {
	http.HandleFunc("/set", set)
	http.HandleFunc("/get", get)
	http.HandleFunc("/del", del)
}

func set(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	item := &memcache.Item{
		Key: "message",
		Value: []byte("hello memcache"),
		// 時間の表記に注意
		Expiration: 30000000000,
	}

	err := memcache.Set(c, item)

	if err == memcache.ErrNotStored {
		fmt.Fprintf(w, "Item is already exist for its key: %s", item.Key)
		return
	} else if err != nil {
		c.Infof("error with memcache.Set() -> %s", err.Error())
		fmt.Fprintf(w, "error with memcache")
		return
	}

	fmt.Fprintf(w, "Item was set to memcache.")
}

func get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	item, err := memcache.Get(c, "message")

	if err != nil {
		c.Infof("error with memcache.Get() -> %s, %v", err.Error(), item)
		fmt.Fprintf(w, "memcache item is not found.")
		return
	}

	fmt.Fprintf(w, "%s", item.Value)
}

func del(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	err := memcache.Delete(c, "message")

	if err != nil {
		c.Infof("error with memcache.Delete() -> %s", err.Error())
		fmt.Fprintf(w, "item could not de deleted.")
		return
	}

	fmt.Fprintf(w, "item was deleted")
}
