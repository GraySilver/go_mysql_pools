package main

import (
	"net/http"
	"fmt"
	_"os"
	_"reflect"
	"io"
	"strings"
	"log"
)


func StartHttpServer(host string,port string,_conn string,maxOpenConns int,maxIdleConns int)  {
	ToInit(_conn,maxOpenConns,maxIdleConns)
	http.HandleFunc("/api/get/",getApi)
	http.HandleFunc("/api/post/",postApi)
	HOST := host + ":"+port
	err := http.ListenAndServe(HOST,nil)
	CheckErr(err)
}


func getApi(w http.ResponseWriter,r *http.Request)  {
	urlQuerys := r.URL.Query()
	if len(urlQuerys) > 0 {
		for k,v := range urlQuerys {
			fmt.Printf("%s=%s\n", k, v[0])
		}
	}
}
func postApi(w http.ResponseWriter,r *http.Request) {

	contextPost := r.PostFormValue("SQL")
	ifSelect := strings.Contains(strings.ToUpper(contextPost),"SELECT")

	if ifSelect{
		resp := ToSelect(contextPost,db)
		result := make(map[string]interface{})
		result["response"] = resp
		result["param"] = r.PostForm
		jsonStr :=MapToJson(result)
		io.WriteString(w, jsonStr)
	}else{
		resp :=ToBegin(contextPost,db)
		result := make(map[string]interface{})
		result["response"] = resp
		result["param"] = r.PostForm
		jsonStr :=MapToJson(result)
		io.WriteString(w,jsonStr)
	}
	log.Printf("[POST]-> params:%s",r.PostForm)
}

