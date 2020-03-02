package controllers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/services"

	"github.com/gorilla/mux"

	amino "github.com/tendermint/go-amino"
)

// StatusController passes requests to its respective service
func StatusController(cdc *amino.Codec, client client.Client, db *db.Database, r *mux.Router) {
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		services.GetStatus(client, db, w, r)
	}).Methods("GET")
}
