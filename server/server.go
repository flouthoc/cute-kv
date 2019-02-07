package main

import (
	"fmt"
	"os"
	"net/http"
	"encoding/gob"
    "io/ioutil"
    "bytes"
)


//core node responsible for storing value;
type CuteNode struct{
	Key string
	Value []byte
}

var cute_kv_global_map map[string]CuteNode

func cute_kv_greet(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, cute_kv_global_map["a"].Value)
	w.Write([]byte("<html><head></head><body>Cute-kv: Is a minimal quick <h4>in-memory</h4> kv store <br>Usage:<br>/get?key=<key> : Will Return Value as response otherwise blank.<br>/set?key=<key>&value=<value> : Will set k:v. Overrides existing key, will return true on sucess.</body></html>"))
}

func cute_kv_set_query(w http.ResponseWriter, r *http.Request){

	formkey := r.FormValue("key")
	formvalue := r.FormValue("value")
	cutenode := CuteNode{Key :formkey, Value: []byte(formvalue)}
	cute_kv_global_map[formkey] = cutenode
	w.Write([]byte("true"))
}


// This function will not be used when running database as in-memory datastore
func cute_kv_set_query_flushable(w http.ResponseWriter, r *http.Request){

	formkey := r.FormValue("key")
	formvalue := r.FormValue("value")
	cutenode := CuteNode{Key :formkey, Value: []byte(formvalue)}
	cute_kv_global_map[formkey] = cutenode
	internal_cute_kv_flush_to_disk()
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


// This function will not be used when running database as in-memory datastore

func internal_cute_kv_load_from_disk(){

	dat, err := ioutil.ReadFile("dat")
	if err != nil{
		fmt.Println("Unable to Read static data")
		fmt.Println("Try Running application with sudo access")
        fmt.Println(err)
	}else{

		buf := bytes.NewBuffer(dat)
		d := gob.NewDecoder(buf)
		err2 := d.Decode(&cute_kv_global_map)
		if err2 != nil{

			fmt.Println("Unable to Decode static data")
			fmt.Println("Try Running application with sudo access")
        	fmt.Println(err2)
		}
	}

}

// This function will not be used when running database as in-memory datastore

func internal_cute_kv_flush_to_disk(){

	b := new(bytes.Buffer)
    e := gob.NewEncoder(b)
    err := e.Encode(cute_kv_global_map)
    if err != nil {
    	fmt.Println("Try Running application with sudo access")
        fmt.Println(err)
        fmt.Println("Something went wrong while serializing the map")
    }else{
    	err2 := ioutil.WriteFile("dat", b.Bytes(), 0777)
    	if err2 != nil{
    		fmt.Println("Try Running application with sudo access")
        	fmt.Println(err2)
        	fmt.Println("Something went wrong while writing map to file")

    	}
    }

}


func main(){

	port := "9992"
	persistance := "false"
	if(len(os.Args) > 1){
		port = os.Args[1]
		if(len(os.Args) > 2){
			persistance = os.Args[2]
		}
	}

	cute_kv_global_map = make(map[string]CuteNode)
	//cute_kv_global_map["a"] = CuteNode{Key: "a", Value: []byte("hello")}

	if persistance == "true" {

		//fill map with contents from disk
		internal_cute_kv_load_from_disk();
		http.HandleFunc("/set", cute_kv_set_query_flushable)

		
	}else{

		http.HandleFunc("/set", cute_kv_set_query)
	}

	http.HandleFunc("/", cute_kv_greet)
	http.HandleFunc("/get", cute_kv_get)
	http.HandleFunc("/health", cute_kv_health)

	fmt.Println("Staring Cute-kv server on port "+port+", visit \n http://127.0.0.1:"+port+"/ \n or http://0.0.0.0:"+port+"/ \n or http://localhost:"+port+"/")

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {

		fmt.Println("Failed to start cute-kv server , reasons could be following")
		fmt.Println("1) Port might me already in use , try some other port")
		fmt.Println("2) Maybe try running server with sudo access")
		fmt.Println("3) Its a just a application listening on a port , worked on my system !")
	}

}