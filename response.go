package wedos

import (
	"encoding/json"
	"net/http"
)

const (
	OK = 1000
)

func (p *Provider) parseResponse(response *http.Response, into any) (respEnv *responseEnvelope, err error) {
	var responseEnvelope responseEnvelope
	if err := json.NewDecoder(response.Body).Decode(&responseEnvelope); err != nil {
		return nil, err
	}

	if responseEnvelope.Response.Code == OK && responseEnvelope.Response.Data != nil && into != nil {
		if err := json.Unmarshal(responseEnvelope.Response.Data, into); err != nil {
			return nil, err
		}
	}

	return &responseEnvelope, nil
}
