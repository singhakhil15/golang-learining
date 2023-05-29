package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ReponseJson struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
	Data      []struct {
		Id        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"data"`
	Support struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}

func main() {
	res, err := http.Get("https://reqres.in/api/users?page=2")

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	responseByte, _ := ioutil.ReadAll(res.Body)

	checkValidate := json.Valid(responseByte)
	if checkValidate {
		fmt.Println(" correct json")
	} else {
		fmt.Println(" Error ")
	}

	var result ReponseJson
	//json.Unmarshal(responseByte, &result)
	if err := json.Unmarshal(responseByte, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(result.Total)
	for _, value := range result.Data {
		fmt.Println(" ===== " + string(value.FirstName))
	}

	fmt.Println(result.Support.Text)
	fmt.Println(result.Support.URL)

}
