package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	model "deserve/model"
)

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	info := model.Info{
		Message:       "Hello Deserve",
		Branches:      "US, Pune, Banglore",
		ContainerName: name,
	}

	json.NewEncoder(w).Encode(info)
}
