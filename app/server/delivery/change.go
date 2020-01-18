package delivery

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/db_project/pkg/messages"
	. "github.com/db_project/pkg/models"
	"github.com/gorilla/mux"
)

func (h *Handler) ChangeThread(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	code := 200

	slugOrID, ok := mux.Vars(r)["slug_or_id"]
	if !ok {
		//
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//
	}

	var threadUpdate ThreadUpdate
	err = json.Unmarshal(bytes, &threadUpdate)
	if err != nil {

	}

	thread, err := h.Service.ChangeThread(threadUpdate, slugOrID)

	if err != nil {
		if err.Error() == messages.ThreadDoesNotExist {
			SetError(w, 404, err.Error())
			return
		}
		log.Println(err.Error())
	}

	answer, _ := json.Marshal(thread)

	w.WriteHeader(code)
	w.Write(answer)
}

func (h *Handler) ChangeUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	code := 200

	nickname, ok := mux.Vars(r)["nickname"]
	if !ok {
		//
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//
	}

	var newUser NewUser
	err = json.Unmarshal(bytes, &newUser)
	if err != nil {

	}

	user, err := h.Service.ChangeUser(newUser, nickname)
	if err != nil {
		if err.Error() == messages.UserAlreadyExists {
			code = 409
		}
		if err.Error() == messages.UserNotFound {
			SetError(w, 404, err.Error())
			return
		}
		log.Println(err.Error())
	}

	answer, _ := json.Marshal(user)

	w.WriteHeader(code)
	w.Write(answer)
}

func (h *Handler) ChangePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	code := 200

	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		//
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		//
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//
	}

	var postUpdate PostUpdate
	err = json.Unmarshal(bytes, &postUpdate)
	if err != nil {

	}

	post, err := h.Service.ChangePost(postUpdate, id)

	if err != nil {
		if err.Error() == messages.PostDoesNotExist {
			SetError(w, 404, err.Error())
			return
		}
		log.Println(err.Error())
	}

	answer, _ := json.Marshal(post)

	w.WriteHeader(code)
	w.Write(answer)
}