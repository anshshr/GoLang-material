package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

)

type Todo struct {
	id          int    `json:"id"`
	uid         int    `json: "uid"`
	description string `json : "descrption"`
	completed   bool   `json:"completed"`
}

func httpPostRequest() {
	url := "https://jsonplaceholder.typicode.com/todos"

	todo := Todo{
		id:          2,
		uid:         23,
		description: "ansh shrivastav",
		completed:   true,
	}

	// convert our data to json using the marshall
	json_data, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("error marshalling the data", err)
		return
	}

	// making a json reader
	jsonreader := strings.NewReader(string(json_data))

	// make an post request now
	res, err := http.Post(url, "application/json", jsonreader)
	defer res.Body.Close()
 
	if err != nil {
		fmt.Println("error pushing the data", err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)

	fmt.Println("status code" , res.StatusCode)

	fmt.Println("response : ", string(data))

}

func httpGetResponse() {
	url := "https://jsonplaceholder.typicode.com/todos/56"
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("error getting the response", err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()



	fmt.Println("data is ", string(data))

}

func httputdateResponse(){
	url := "https://jsonplaceholder.typicode.com/todos/1"

	todo := Todo{
		id:          2,
		uid:         234,
		description: "ansh shrivastav is a golang developer",
		completed:   true,
	}

	json_data , err := json.Marshal(todo);
	

	if err != nil{
		fmt.Println("error getting the data" , err);
		return;
	}

	// convert json data to string
	json_string  := string(json_data);

	// make a reader
	json_reader := strings.NewReader(json_string)

	// create a put or patch request to put the data on server

	req , err := http.NewRequest(http.MethodPut , url, json_reader);

	if err != nil{
		fmt.Println("error putting  the data" , err);
		return;
	}

	// set the headers
	req.Header.Set("Content-Type" , "application/json")

	// send the request

	// creating the client
	client := http.Client{};
	// make an request on clinet
	res , err := client.Do(req);

	if err != nil{
		fmt.Println("error getting the response " , err)
		return;
	}

	defer res.Body.Close()

	byte_data , err := ioutil.ReadAll(res.Body)

	if err != nil{
		fmt.Println("error getting the data ", err);
		return;
	}

	fmt.Println("the data is" , string(byte_data))

	fmt.Println(res.StatusCode)



}

func httpDeleteRequest(){
	url := "https://jsonplaceholder.typicode.com/todos/2"

	todo := Todo{
		id:          2,
		uid:         234,
		description: "ansh shrivastav is a golang developer",
		completed:   true,
	}
	
	// comverting to json data
	json_data , err := json.Marshal(todo);
	
	if err != nil{
		fmt.Println("error getting the response" , err);
		return;
	}

	json_string := string(json_data);

	json_reader := strings.NewReader(json_string);


	req, err := http.NewRequest(http.MethodDelete , url , json_reader)

	if(err != nil){
		fmt.Println("error getting the req" , err);
		return;
	}
	req.Header.Set("Content-Type" , "application/json");

	client := http.Client{};
	res , err := client.Do(req)

	if err != nil{
		fmt.Println("error getting the response", err);
		return;
	}

	defer res.Body.Close()

	byte_data , err := ioutil.ReadAll(res.Body)

	fmt.Println("the deleted data is " , string(byte_data))
	fmt.Println("the status code is" , res.StatusCode)

}

func main() {
	fmt.Println("learning various crud operations")
	// httpGetResponse()
	// httpPostRequest()
	// httputdateResponse()
	httpDeleteRequest()
}
