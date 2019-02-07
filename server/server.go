package main

import (
	"fmt"
	"net/http"
)


//core node responsible for storing value;

type CuteNode struct{
	Key string
	Value []byte
}



var cute_kv_global_map map[string]CuteNode

func cute_kv_greet(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, cute_kv_global_map["a"].Value)
	w.Write([]byte("<html><head></head><body>Cute-kv: Is a minimal quick persistant kv store 0<br>Usage:<br>/get?key=<key> : Will Return Value as response otherwise blank.<br>/set?key=<key>&value=<value> : Will set k:v. Overrides existing key, will return true on sucess.</body></html>"))
}

func cute_kv_set_query(w http.ResponseWriter, r *http.Request){

	formkey := r.FormValue("key")
	formvalue := r.FormValue("value")
	cutenode := CuteNode{Key :formkey, Value: []byte(formvalue)}
	cute_kv_global_map[formkey] = cutenode
	w.Write([]byte("true"))
}

func cute_kv_get(w http.ResponseWriter, r *http.Request){

	cutenode, ok := cute_kv_global_map[r.FormValue("key")]

	if ok{

		w.Write(cutenode.Value)

	}else{

		w.Write([]byte("nil"))
	}

}

func cute_kv_health(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("healthy"))
}


/*func internal_cute_kv_load_from_disk(){

}*/

func internal_cute_kv_flush_to_disk(){

}


func main(){

	cute_kv_global_map = make(map[string]CuteNode)
	//cute_kv_global_map["a"] = CuteNode{Key: "a", Value: []byte("hello")}

	http.HandleFunc("/", cute_kv_greet)
	http.HandleFunc("/set", cute_kv_set_query)
	http.HandleFunc("/get", cute_kv_get)
	http.HandleFunc("/health", cute_kv_health)

	fmt.Println("Staring Cute-kv server on port 8080, visit \n http://127.0.0.1:8080/ \n or http://0.0.0.0:8080/ \n or http://localhost:8080/")

	http.ListenAndServe(":8080", nil)

}