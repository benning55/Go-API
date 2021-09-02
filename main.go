package main

import (
	"fmt"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		`<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		<body>
			<h1>This is our server Benning</h1>
		</body>
		</html>`)
}

func myWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "thia ia welcome")
}

func myLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is login")
}

func main() {
	http.HandleFunc("/login", myLogin)
	http.HandleFunc("/welcome", myWelcome)
	fmt.Println("Listening on port 8080....")
	http.ListenAndServe("localhost:8080", nil)
}
