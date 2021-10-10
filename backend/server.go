package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Input
type skillSearch struct {
	Skill string 'json:"skill"'
}

// Output
type listItemData struct {
	Name string 'json:"name"'
	Price int64 'json:"price"'
	Remote bool 'json:"remote"'
	Location string 'json:"location"'
}

type listResponseData struct {
	Results []listItemData 'json:"results"'
}

// Process Request for index results
func process(requestData skillSearch) (listResponseData) {
	var listResponse listResponseData
	return listResponse
}

// Send Response
func getResponse(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var skill skillSearch
	var response listResponseData

	decode.Decode(&skill)

	response = process(skill)
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func main(){
	http.HandleFunc("/", getResponse)
	http.ListenAndServe(":8090", nil)
}
