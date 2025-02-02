package main

// air package - similar to nodemon in node js to continously run the server(fresh , reflex and compiledaemon are other alternatives)

// here we are using the "mux/gorilla" which is simila to express js so we create route we can use http 
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("building a go server")

	route := mux.NewRouter()		// insatll gorilla/mux for dyanmic routing so we can implement the dynamic routing in golang

	// api routes
	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //func is similar to fat arrow function in js
		fmt.Fprintf(w, "started the sever and it is runnning daam fine") // to print or show anythig in the webpage
		// fmt.Println("we actually started the server ans")
	})

	route.HandleFunc("/url_shortner", func(w http.ResponseWriter, r *http.Request) { // customised routes
		fmt.Fprintf(w, "url page for the app is running live now and i am doing it")
	})

	route.HandleFunc("/url_shortner/{encoded_url}", func(w http.ResponseWriter, r *http.Request) { // use of he query string(dyanmc routing using the gorilla mux) in golang similar to node js

		vars := mux.Vars(r);
		endcoded_url := vars["encoded_url"]
		// fmt.Fprintf(w, endcoded_url)

		fmt.Println(endcoded_url)

		http.Redirect(w ,r , "https://www.geeksforgeeks.org/how-to-build-a-simple-web-server-with-golang/" , http.StatusFound) // we can pass the encoded url to do the same thing 
		
	})

	port := ":3000"

	log.Fatal(http.ListenAndServe(port, route))


}
