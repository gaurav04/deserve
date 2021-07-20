package handlers

import (
	model "deserve/model"
	"encoding/json"
	"net/http"
	"os"
)

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	info := model.Info{
		Message:       "Hello Talentica",
		Branches:      "Pune",
		Employee:      450,
		ContainerName: name,
	}

	json.NewEncoder(w).Encode(info)
}
