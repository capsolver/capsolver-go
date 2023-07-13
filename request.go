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
		ApiHost = "https://api.capsolver.com"
	}
}

func request(uri string, solverRequest *capSolverRequest) (*capSolverResponse, error) {

	b, _ := json.Marshal(solverRequest)
	resp, err := http.Post(fmt.Sprintf("%s%s", ApiHost, uri), "json/application", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, _ = io.ReadAll(resp.Body)
	capResponse := &capSolverResponse{}
	capResponse.Solution = &solution{}
	json.Unmarshal(b, capResponse)
	return capResponse, nil
}
