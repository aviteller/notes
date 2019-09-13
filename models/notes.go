package models

import (
	"fmt"
	"strconv"

	u "../utils"
	"github.com/sonyarouje/simdb/db"
)

type Note struct {
	NoteID string `json:"ID"`
	Title  string `json:"title"`

	Message string `json:"message"`
}

func (c Note) ID() (jsonField string, value interface{}) {

	value = c.NoteID

	jsonField = "ID"

	return

}

func (note *Note) Validate() (map[string]interface{}, bool) {

	if note.Title == "" || note.Message == "" {
		return u.Message(false, "Please enter all fields"), false
	}

	return u.Message(true, "success"), true
}

func (note *Note) Create() map[string]interface{} {
	if res, ok := note.Validate(); !ok {
		return res
	}

	driver, err := db.New("dbs")

	if err != nil {

		panic(err)

	}

	note.NoteID = GetLastNoteID()

	err = driver.Insert(note)

	if err != nil {
		panic(err)
	}

	res := u.Message(true, "success")
	res["note"] = note
	return res
}

//GetLastNoteID gets all the meetups
func GetLastNoteID() string {
	notes := make([]*Note, 0)
	driver, err := db.New("dbs")

	if err != nil {

		panic(err)

	}

	err = driver.Open(Note{}).Get().AsEntity(&notes)

	if err != nil {

		panic(err)

	}
	ID := "1"

	if len(notes) > 0 {

		for i := len(notes) - 1; i >= 0; i-- {

			lastid, err := strconv.Atoi(notes[i].NoteID)
			if err != nil {

				fmt.Println(err)
			}

			lastid = lastid + 1
			ID = strconv.Itoa(int(lastid))
			break
		}
	}

	return ID
}

//GetBirthdays gets all the meetups
func GetNotes() []*Note {
	notes := make([]*Note, 0)
	driver, err := db.New("dbs")

	if err != nil {

		panic(err)

	}

	err = driver.Open(Note{}).Get().AsEntity(&notes)

	if err != nil {

		panic(err)

	}

	return notes
}

func DeleteNote(id string) map[string]interface{} {

	driver, err := db.New("dbs")

	if err != nil {

		panic(err)

	}

	toDel := Note{

		NoteID: id,
	}

	err = driver.Delete(toDel)

	if err != nil {

		panic(err)
	}
	res := u.Message(true, "success")

	return res
}

func (note *Note) UpdateNote(id string) map[string]interface{} {
	if res, ok := note.Validate(); !ok {
		return res
	}

	driver, err := db.New("dbs")

	if err != nil {

		panic(err)

	}

	err = driver.Update(note)
	if err != nil {
		panic(err)
	}

	res := u.Message(true, "success")

	return res
}
