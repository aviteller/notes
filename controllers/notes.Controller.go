package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
	u "../utils"
	"github.com/gorilla/mux"
)

//CreateNote is an action the creates a note and then sends the data tot eh model
var CreateNote = func(w http.ResponseWriter, r *http.Request) {
	//user := r.Context().Value("user").(uint)
	note := &models.Note{}

	err := json.NewDecoder(r.Body).Decode(note)
	if err != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	//contact.UserId = user
	res := note.Create()
	u.Respond(w, res)
}

//GetNotes gets all the birthdays
var GetNotes = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetNotes()

	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)

}

var UpdateNote = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	note := &models.Note{}

	err2 := json.NewDecoder(r.Body).Decode(note)
	if err2 != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	res := note.UpdateNote(id)
	res["data"] = id
	u.Respond(w, res)
}

var DeleteNote = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	res := models.DeleteNote(id)
	res["data"] = id
	u.Respond(w, res)
}
