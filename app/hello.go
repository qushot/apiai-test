package hello

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/qushot/apiai-test/model"
)

func init() {
	http.HandleFunc("/hello", helloHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "hello")

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	var req model.APIAIRequest
	err := json.Unmarshal(buf.Bytes(), &req)
	if err != nil {
		log.Errorf(ctx, "%v", err)
	}

	log.Infof(ctx, "BODY=%v", req)

	w.Header().Set("Content-type", "application/json")
	res := model.APIAIResponse{}
	res.Speech = "I'm API.AI!"
	res.DisplayText = "I'm API.AI!"
	data, err := json.Marshal(res)
	if err != nil {
		log.Errorf(ctx, "%v", err)
	}
	w.Write(data)

	fmt.Fprint(w, "")
}
