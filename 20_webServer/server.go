package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
1. Golang provides a built-in HTTP package that contains utilities for quickly creating a web or file server.

We will create a web server that can accept a GET request and serve a response.
We’ll then make the web server respond to a POST request coming from a form submission, such as a contact form.
*/
//3.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

/* 6th code
First, the function has to call ParseForm() to parse the raw query and update r.PostForm and
r.Form. This will allow us to access the name and address values via the r.FormValue method.

At the end of the function, we write both values to the ResponseWriter using fmt.Fprintf.

*/

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	mobile := r.FormValue("mobile")
	city := r.FormValue("city")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Mobile = %s\n", mobile)
	fmt.Fprintf(w, "Mobile = %s\n", city)
}

func main() {
	fmt.Printf("Starting server at port 8080\n")

	//2. the next step is to create a web server
	//we’ll create a web server that is actually served on port 8080 and can respond to incoming GET requests.

	/* Note that the port parameter needs to be passed as a string prepended by colon punctuation.
	The second parameter accepts a handler to configure the server for HTTP.
	However, this isn’t important for this tutorial, so we can safely pass nil as the second argument.
	*/
	//2
	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello!")
	// })

	//4.
	fileServer := http.FileServer(http.Dir("./html")) // New code
	http.Handle("/", fileServer)                      // New code

	//3.
	http.HandleFunc("/hello", helloHandler) // Update this line of code

	//6th
	http.HandleFunc("/contact", formHandler)

	//1.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	/* 3. the server can start, but it still doesn’t know how to handle requests.
		We need to pass handlers to the server so it knows how to respond to incoming
		requests and which requests to accept.

		We’ll use the HandlerFunc function to add route handlers to the web server.
		The first argument accepts the path it needs to listen for /hello.
		Here, you tell the server to listen for any incoming requests for http://localhost:8080/hello.
		The second argument accepts a function that holds the business logic to correctly respond to the request.

		By default, this function accepts a ResponseWriter to send a response back and
		a Request object that provides more information about the request itself.

		4. Add basic security to routes
	It goes without saying that security is important. Let’s explore some basic strategies
	to enhance the security of your Go web server.

	Before we do, we should take a moment to increase the readability of our code.
	Let’s create the helloHandler function, which holds all the logic related to the /hello request.

	5.
	we’ll create a simple file server to host static files. This will be a very simple addition to the web server.
	To serve the static folder, you’ll have to add two lines of code to server.go.
	The first line of code creates the file server object using the FileServer function.
	This function accepts a path in the http.Dir type. Therefore, we have to convert the string path
	“./static” to an http.Dir path type.

	6.
	create a form submission using POST request
	create contact.html file
	The form will take three inputs: name, mobile and city

	The next step is to create the handler to accept the /contact request
	*/

}
