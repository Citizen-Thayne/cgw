package api

import (
	"bytes"
	"cgw/pkg/game"
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func eventStream(c *gin.Context) {
	w := c.Writer
	r := c.Request

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Create a channel to send data
	eventChan := make(chan int, 10)
	GetInstance().Game.Subscribe(eventChan)

	// Create a context for handling client disconnection
	_, cancel := context.WithCancel(r.Context())
	defer cancel()

	// Send data to the client
	for event := range eventChan {
		switch event {
		case game.NewGeneration:
			templ, err := template.ParseFiles("templates/game.go.html")

			if err != nil {
				panic(err)
			}

			fmt.Fprint(w, "event: NewGeneration\ndata: ")
			var buff bytes.Buffer
			err = templ.Execute(&buff, gin.H{"game": GetInstance().Game})

			data := buff.Bytes()
			data = bytes.ReplaceAll(data, []byte{'\n'}, []byte{})
			w.Write(data)

			if err != nil {
				panic(err)
			}

			fmt.Fprint(w, "\n\n")

			w.(http.Flusher).Flush()

		case game.GamePaused:
			templ, err := template.ParseFiles("templates/controls_play.go.html")

			if err != nil {
				panic(err)
			}

			fmt.Fprint(w, "event: PlayButton\ndata: ")
			var buff bytes.Buffer
			err = templ.Execute(&buff, gin.H{"isPlaying": false})

			data := buff.Bytes()
			data = bytes.ReplaceAll(data, []byte{'\n'}, []byte{})
			w.Write(data)

			if err != nil {
				panic(err)
			}

			fmt.Fprint(w, "\n\n")

			w.(http.Flusher).Flush()
		case game.GameResumed:
			templ, err := template.ParseFiles("templates/controls_play.go.html")

			if err != nil {
				panic(err)
			}

			fmt.Fprint(w, "event: PlayButton\ndata: ")
			var buff bytes.Buffer
			err = templ.Execute(&buff, gin.H{"isPlaying": true})

			data := buff.Bytes()
			data = bytes.ReplaceAll(data, []byte{'\n'}, []byte{})
			w.Write(data)

			if err != nil {
				panic(err)
			}

			fmt.Fprint(w, "\n\n")

			w.(http.Flusher).Flush()
		}
	}
}
