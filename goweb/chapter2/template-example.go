package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template_example.tmpl")
	if err != nil {
		fmt.Println("template pasefile failed ,err:", err)
		return
	}
	name := "sbzhc"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", helloHandleFunc)
	http.ListenAndServe(":8086", nil)
}
