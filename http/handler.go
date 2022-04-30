package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Uzama/most_used_words/processor"
)

type Body struct {
	Text string `json:"text"`
}

func TextHandler(w http.ResponseWriter, r *http.Request) {
	// read from the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	defer r.Body.Close()

	b := &Body{}

	// decode the body
	err = json.Unmarshal(body, b)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// check whether given text is not empty
	if b.Text == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("text cannot be empty"))
		return
	}

	// process the text
	words := processor.TextProcessor(b.Text)

	payload, _ := json.Marshal(words)

	// respond the result
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
