package capsolver_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func init() {
	if ApiHost == "" {
		ApiHost = DEFAULT_API_HOST
	}
}

func request(uri string, solverRequest *capSolverRequest) (*CapSolverResponse, error) {
	b, _ := json.Marshal(solverRequest)
	resp, err := http.Post(fmt.Sprintf("%s%s", ApiHost, uri), "application/json", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	b, _ = io.ReadAll(resp.Body)
	capResponse := &CapSolverResponse{}
	capResponse.Solution = &solution{}
	err = json.Unmarshal(b, capResponse)
	if err != nil {
		return nil, err
	}
	return capResponse, nil
}
