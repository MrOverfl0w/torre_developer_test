package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Jeffail/gabs"
)

// Input
type skillSearch struct {
	Skill string `json:"skill"`
}

// Output
type listItemData struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Remote   bool    `json:"remote"`
	Location string  `json:"location"`
}

type listResponseData struct {
	Results []listItemData `json:"results"`
}

// Process Request for index results
func process(requestData skillSearch) listResponseData {
	var listResponse listResponseData

	usersjson, err := http.Post("https://search.torre.co/people/_search/?offset=0&size=1000", "", nil)
	if err != nil {
		log.Fatalf("Error in API: %v", err)
	}
	defer usersjson.Body.Close()
	jsonResponse, err := ioutil.ReadAll(usersjson.Body)
	if err != nil {
		log.Fatalln(err)
	}

	jsonParsed, _ := gabs.ParseJSON([]byte(jsonResponse))
	children, _ := jsonParsed.Path("results").Children()
	//Children Array from Results
	for _, child := range children {
		mentor := false
		childrenOpen, _ := child.Path("openTo").Children()
		for _, childOpen := range childrenOpen {
			if childOpen.Data().(string) == "mentoring" {
				mentor = true
			}
		}

		if mentor {
			var weight float64
			weight = 0
			children2, _ := child.Path("skills").Children()
			//Children Array from Skills
			for _, child2 := range children2 {
				if child2.Path(("name")).Data() == requestData.Skill {
					weight = child2.Path(("weight")).Data().(float64)
				}
			}
			if weight > 0 {
				var itemDataResponse listItemData
				itemDataResponse.Name = child.Path("name").Data().(string)
				itemDataResponse.Price = child.Path("compensations").Path("employee").Path("minHourlyUSD").Data().(float64)
				itemDataResponse.Remote = child.Path("remoter").Data().(bool)
				itemDataResponse.Location = child.Path("locationName").Data().(string)

				listResponse.Results = append(listResponse.Results, itemDataResponse)
			}
		}

	}

	return listResponse
}

// Send Response
func getResponse(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var skill skillSearch
	var response listResponseData

	decoder.Decode(&skill)
	log.Println(skill.Skill)

	response = process(skill)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", getResponse)
	http.ListenAndServe(":8090", nil)
}
