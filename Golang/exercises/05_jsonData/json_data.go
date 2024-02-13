package jsondata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type GetAllUserResponse struct {
	Page       int           `json:"page"`
	PerPage    int           `json:"per_page"`
	Total      int           `json:"total"`
	TotalPages int           `json:"total_pages"`
	Data       []interface{} `json:"data"`
}

type CreateUserRequest struct {
	Name string
	Job  string
}

type CreateUserResponse struct {
	Name      string `json:"name"`
	Job       string `json:"job"`
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

func GetAllUser() {
	fmt.Println("Executing GetUser Function")

	//http get request
	httpResponse, err := http.Get("https://reqres.in/api/users/")
	if err != nil {
		fmt.Println("Error in making api call", err)
		os.Exit(1)
	}

	//parse the response
	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println("Error while reading response data", err)
		os.Exit(1)
	}
	// var x interface{}
	var response GetAllUserResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		fmt.Println("Error while unmarshalling response data", err)
	} else {
		//Print the response data
		fmt.Printf("Page: %d\n", response.Page)
		fmt.Printf("Per Page: %d\n", response.PerPage)
		fmt.Printf("Total: %d\n", response.Total)
		fmt.Printf("Total Pages: %d\n", response.TotalPages)

		// Print user data of type []interface{}
		for _, userData := range response.Data {
			// Since Data is of type []interface{}, we need to assert its type before using it
			if user, ok := userData.(map[string]interface{}); ok {
				fmt.Printf("User ID: %v, Email: %v, First Name: %v, Last Name: %v, Avatar: %v\n", user["id"], user["email"], user["first_name"], user["last_name"], user["avatar"])
			} else {
				fmt.Println("Error: Unable to parse user data.")
			}
		}
	}
}

func CreateUser() {
	fmt.Println("Executing CreateUser Function")

	//Declarations
	user := CreateUserRequest{Name: "Quizra", Job: "Software Engineer"}
	postURL := "https://reqres.in/api/users"

	//cannot convert CreateUserRequest which is of type struct to type []byte
	//so using json.Marshal to convert it to []byte
	resp, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error while marshalling user data", err)
	}

	//http post request
	httpResponse, err := http.Post(postURL, "application/json", bytes.NewBuffer(resp))
	if err != nil {
		fmt.Println("Error while posting the user data", err)
	}

	//parse the response
	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println("Error while reading response data", err)
	}
	var response CreateUserResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		fmt.Println("Error while unmarshalling response data", err)
	} else {
		//Print the response data
		fmt.Printf("Data: %s\n", response)
		fmt.Printf("Name: %s\n", response.Name)
		fmt.Printf("Job: %s\n", response.Job)
		fmt.Printf("Id: %s\n", response.Id)
		fmt.Printf("Created At: %s\n", response.CreatedAt)
	}
}
