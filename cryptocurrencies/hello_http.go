package cryptocurrencies

import (
	"encoding/json"
	"net/http"

	"github.com/MalukiMuthusi/cryptocurrencies/fauna"
	"github.com/MalukiMuthusi/cryptocurrencies/logger"
	"github.com/MalukiMuthusi/cryptocurrencies/models"
)

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	// initialize the storage service
	store, err := fauna.New()
	if err != nil {
		logger.Log.Fatal("failed to initialize database ", err)
	}

	basicError := models.BasicError{
		Code:    "SERVER_ERROR",
		Message: "failed to get cryptocurrencies",
	}

	l, err := store.List(r.Context())
	if err != nil {

		b, err := json.Marshal(basicError)
		if err != nil {
			return
		}

		w.Write(b)
		return
	}

	b, err := json.Marshal(l)
	if err != nil {

		b, err := json.Marshal(basicError)
		if err != nil {
			return
		}

		w.Write(b)
		return
	}

	w.Write(b)
}
