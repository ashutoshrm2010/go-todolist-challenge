package app

import (
	"net/http"
	"github.com/PeakActivity/go-todolist-challenge/app/lib"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"fmt"
	"errors"
	"github.com/pborman/uuid"
	"gopkg.in/mgo.v2/bson"
)

// Message ...
type Message struct {
	Message string `json:"message"`
}

func InsertItem(w http.ResponseWriter, r *http.Request) {
	res := lib.Response{ResponseWriter: w}
	decoder := json.NewDecoder(r.Body)
	var data map[string]interface{}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, errors.New("please provide some entry").Error(), http.StatusUnauthorized)
		return
	}
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	data["_id"] = uuid.New()
	c := session.DB("peakActivity").C("items")
	err = c.Insert(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	res.SendOK(data)
}
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	res := lib.Response{ResponseWriter: w}
	decoder := json.NewDecoder(r.Body)
	var data map[string]interface{}
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("error ", err)
		http.Error(w, errors.New("please provide some entry").Error(), http.StatusUnauthorized)
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("peakActivity").C("items")
	data["_id"] = uuid.New()
	where := bson.M{"_id":data["_id"].(string)}
	update := bson.M{"name":data["name"].(string)}
	err = c.Update(where, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	res.SendCreated(data)
}
func ListItem(w http.ResponseWriter, r *http.Request) {
	res := lib.Response{ResponseWriter: w}
	var data []map[string]interface{}
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("peakActivity").C("items")
	err = c.Find(nil).All(&data)
	if err != nil {
		if err.Error() == "not found" {
			m := Message{"No items found"}
			res.SendOK(m)
			return
		} else {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

	}
	if len(data) > 0 {
		res.SendOK(data)

	} else {
		m := Message{"No items found"}
		res.SendOK(m)
	}
}
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	res := lib.Response{ResponseWriter: w}
	decoder := json.NewDecoder(r.Body)
	var data map[string]interface{}
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("error ", err)
		http.Error(w, errors.New("please provide some entry").Error(), http.StatusUnauthorized)
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("peakActivity").C("items")
	err = c.Remove(bson.M{"_id":data["_id"].(string)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	res.SendOK(data)
}
