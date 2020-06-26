package main

import (
    "fmt"
	"log"
	"net/http"
	"encoding/json"
)


func homePage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Homepage")
}


type Article struct{
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Author string `json:"Author"`
}

type Articles []Article

func allArticle(w http.ResponseWriter,r *http.Request){
	articles := Articles{
		Article{Title : "Test Title",Desc : "Test Desc",Author : "Test Author"},
	}
	fmt.Println("All Articles"	)
	json.NewEncoder(w).Encode(articles)
	
}



// Form Handle Function

func formHandle(w http.ResponseWriter , r *http.Request){
	if err := r.ParseForm(); err!=nil{
		fmt.Fprintf(w ,"Parse Form Error : %v",err)
		return
	}
	fmt.Fprintf(w,"POST REQUEST SUCCESS \n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	
	fmt.Fprintf(w,"Name %s \n",name)
	fmt.Fprintf(w,"Address %s \n",address)
	

}


func handleRequests(){
	http.HandleFunc("/homepage",homePage)
	http.HandleFunc("/articles",allArticle)
	http.HandleFunc("/form",formHandle)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func main() {
	
	fileServer := http.FileServer(http.Dir("./static"))  //HTML File Folder
	http.Handle("/",fileServer)
	handleRequests()
    fmt.Println("Hello World")
}

https://github.com/nic-delhi/AarogyaSetu_Android
https://dzone.com/articles/scaling-microservices-the-challenges-and-solutions :)
