package api

import (
	"encoding/json"
	"github.com/evilsocket/islazy/log"
	"io"
	"net/http"
)

// PeerGetData GET /api/v1/data
func (api *API) PeerGetData(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, api.Client.Data())
}

// PeerSetData POST /api/v1/data
func (api *API) PeerSetData(w http.ResponseWriter, r *http.Request) {
	var newData map[string]interface{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	log.Debug("%s", body)

	if err = json.Unmarshal(body, &newData); err != nil {
		ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	JSON(w, http.StatusOK, api.Client.SetData(newData))
}
