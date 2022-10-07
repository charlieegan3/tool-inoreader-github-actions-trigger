package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/charlieegan3/tool-inoreader-github-actions-trigger/pkg/api"
)

func BuildGetHandler(targets map[string]api.Target) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var err error
		vars := mux.Vars(request)

		target, ok := vars["target"]
		if !ok || target == "" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		loadedTarget, ok := targets[target]
		if !ok {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		var payload api.InoreaderPost
		err = json.NewDecoder(request.Body).Decode(&payload)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(payload.Items) == 0 {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		forwardedData := map[string]interface{}{
			"event_type": loadedTarget.EventType,
			"client_payload": map[string]interface{}{
				"url":   payload.Items[0].Canonical[0].Href,
				"title": payload.Items[0].Title,
			},
		}

		bodyJSON, err := json.Marshal(forwardedData)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}

		client := http.Client{Timeout: 10 * time.Second}
		req, err := http.NewRequest("POST", loadedTarget.URL, bytes.NewReader(bodyJSON))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+loadedTarget.Token)

		resp, err := client.Do(req)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}

		if resp.StatusCode >= 400 {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("unexpected response from github: %d", resp.StatusCode)))
			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}
