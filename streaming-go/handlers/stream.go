package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"streaming-go/model"
	"time"
)

func StreamTime(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Optional for CORS

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.Context().Done():
			fmt.Println("Client disconnected")
			return
		case t := <-ticker.C:
			event := model.Event{
				ID:      fmt.Sprintf("%d", t.Unix()),
				Message: "Current server time",
				Time:    time.Now().String(),
			}

			jsonData, err := json.Marshal(event)
			if err != nil {
				continue
			}

			fmt.Fprintf(w, "data: %s\n\n", string(jsonData))
			flusher.Flush()
		}
	}
}
