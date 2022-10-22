package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)


func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<html>"+
		"<head>"+
		"<title>Home Web-server</title>"+
		"</head>"+
		"<body>"+
		"<h1>Hi I run this server on premise in Seattle, Washington, USA! This is me :) -> </h1>"+
		"<p> For the sake of using golang I will tell you the time: "+ time.Now().String() +"</p>"+
		"<a href=\"https://hamdaan-rails-personal.herokuapp.com/\">Visit Me Here!</a>" +
		"</body>"+
		"</html>")
    })


    fmt.Printf("Started server at port 3001\n")
    if err := http.ListenAndServe(":3001", nil); err != nil {
        log.Fatal(err)
    }

}
