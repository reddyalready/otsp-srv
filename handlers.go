package main

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
	"github.com/reddyalready/otsp-srv/db"
	"github.com/reddyalready/otsp-srv/model"
)

var projects = model.Projects{
	model.Project{Name: "Create a side projects website"},
	model.Project{Name: "Fake side project"},
}

func ProjectList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.ProjectList())
}

func ProjectGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if projectID, err := strconv.Atoi(params["projectId"]); err == nil {

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(projects[projectID]); err != nil {
			panic(err)
		}
	}
}
