package lambdaprox

import (
	"io/ioutil"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/corpix/uarand"
)

func init() {
	functions.HTTP("LambdaProx", LambdaProx)
}

func LambdaProx(w http.ResponseWriter, r *http.Request) {
	// Get the "url" query parameter from the request
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
		return
	}

	// Fetch data from the provided URL
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	req.Header.Set("User-Agent", uarand.GetRandom())
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the fetched data directly to the response body
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
