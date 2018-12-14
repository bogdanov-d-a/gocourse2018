package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVideo(t *testing.T) {
	w := httptest.NewRecorder()
	videoImpl(w, "d290f1ee-6c54-4b01-90e6-d701748f0851")
	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusOK)
	}

	jsonString, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	item := VideoData{}
	if err = json.Unmarshal(jsonString, &item); err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}
}
