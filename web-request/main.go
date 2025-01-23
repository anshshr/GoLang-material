package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("now  the main part is starting - web request")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("error fetching the response", err)
		return
	}
	if res.StatusCode == 200 {
		fmt.Println("response is :", res) // the data return is cannot be readable we had to convert it into bytes and then to string to read the data
	}

	// reading the data from the file
	content, _ := io.ReadAll(res.Body) // now it contain the data in the form of bytes
	fmt.Println(content)

	fmt.Println("actual data is :", string(content)) // conversion of bytes to string

	// leraning about the urls (go lang provides a url package to exatrct and modify different parts of the url)

	myURL := "https://example.com/to/your/path?key1=value1,key2=value2";

	parsed_url , err := url.Parse(myURL);

	if err != nil{
		fmt.Println("error fething the url" , err);
		return;
	}

	fmt.Println(parsed_url.Path) // path of the url
	fmt.Println(parsed_url.Host) // host name of the url
	fmt.Println(parsed_url.RawQuery) // query string of the url
	fmt.Println(parsed_url.Scheme)	// scheme of the url


	// modifying the url
	parsed_url.Path = "newpath";
	parsed_url.Scheme = "www"

	new_string := parsed_url.String()

	fmt.Println(new_string)

}
