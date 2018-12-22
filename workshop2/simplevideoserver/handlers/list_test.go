package handlers

import (
	"encoding/json"
	"github.com/bogdanov-d-a/gocourse2018/workshop2/simplevideoserver/database"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {
	w := httptest.NewRecorder()
	List(make([]database.VideoData, 0), w)
	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusOK)
	}

	jsonString, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	items := make([]ListData, 10)
	if err = json.Unmarshal(jsonString, &items); err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}
}
