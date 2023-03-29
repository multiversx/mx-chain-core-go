package mock

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/websocket"
	logger "github.com/multiversx/mx-chain-logger-go"
)

var upgrader = websocket.Upgrader{} // use default options
var log = logger.GetOrCreate("websocketOutportDriver/mock")

type httpTestEchoHandler struct{}

// ServeHTTP -
func (handler *httpTestEchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/echo":
		echo(w, r)
	default:
		http.NotFound(w, r)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, errUpgrade := upgrader.Upgrade(w, r, nil)
	if errUpgrade != nil {
		return
	}
	defer func() {
		err := c.Close()
		log.LogIfError(err)
	}()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}
		stringMessage := string(message)

		log.Info("received on server", "message", stringMessage)
		err = c.WriteMessage(mt, []byte("ECHO: "+stringMessage))
		if err != nil {
			break
		}
	}
}

// NewHttpTestEchoHandler -
func NewHttpTestEchoHandler() *httptest.Server {
	return httptest.NewServer(&httpTestEchoHandler{})
}
