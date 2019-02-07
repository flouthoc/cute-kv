package main

import (

	"os"
	"fmt"
	"net/http"
	"io/ioutil"
)

func clientWatch(host string, key string){

	lastoutput := ""
	for true {
		val := clientGet(host, key)
		if val != lastoutput {
			lastoutput = val
			fmt.Println(lastoutput)
		}
	}

}


func clientGet(host string, key string) string{

	req, err := http.NewRequest("GET", "http://"+host+"/get", nil)
    if err != nil {
        return "nil"
    }

    q := req.URL.Query()
    q.Add("key", key)
    req.URL.RawQuery = q.Encode()

    resp, err2 := http.Get(req.URL.String())
	if err2 != nil {
		return "nil"
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	content := string(body[:])
	return content

}


func clientSet(host string, key string, value string) bool {
	req, err := http.NewRequest("GET", "http://"+host+"/set", nil)
    if err != nil {
        return false
    }

    q := req.URL.Query()
    q.Add("key", key)
    q.Add("value", value)
    req.URL.RawQuery = q.Encode()

    resp, err2 := http.Get(req.URL.String())
	if err2 != nil {
		return false
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	content := string(body[:])
	fmt.Println(content)
	return true
}

func verifyHost(host string) bool {

	resp, err := http.Get("http://"+host+"/health")
	if err != nil {
		return false
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	content := string(body[:])

	if content == "healthy" {

		return true

	}else{

		return false
	}

	return false

}


func main(){

	usage := "cute-kv client manual page \n --- \n Gets value for specified key \n\n  Syntax : get <hostip:port> <key> \n  Example : ./cliclient get 127.0.0.1:9992 samplekey \n\n --- \n Sets value for specified key \n\n  Syntax : set <hostip:port> <key> <value> \n  Example : ./cliclient set 127.0.0.1:9992 samplekey samplevalue \n\n --- \n Watches key for change in a blocking mode \n\n  Syntax : watch <hostip:port> <key> \n  Example : ./cliclient watch 127.0.0.1:9992 samplekey\n\n ---";
	//fmt.Println(usage)

	action := os.Args[1]
	argslen := len(os.Args)

	if(argslen < 3){

		fmt.Println(usage)

	}else{

		host := os.Args[2]

		if action == "set" {

			if verifyHost(host) == true {

				if(argslen == 5){

					key := os.Args[3]
					value := os.Args[4]

					if clientSet(host, key, value){
						fmt.Println("Done.")
					}else{
						fmt.Println("Something went wrong.")
					}

				}else{

					fmt.Println("Please see correct usage of set command")
					fmt.Println("Example: ./cliclient set 127.0.0.1:9992 samplekey samplevalue")
				}

				

			}else{

				fmt.Println("Looks like server is not running on "+host)
				fmt.Println("Try running server on "+host)
				fmt.Println("Example: ./cliclient set 127.0.0.1:9992 samplekey samplevalue")
			}


		}else if action == "get" {

			if verifyHost(host) == true {

				if(argslen == 4){

					key := os.Args[3]
					fmt.Println(clientGet(host, key))
					

				}else{

					fmt.Println("Please see correct usage of get command")
					fmt.Println("Example: ./cliclient get 127.0.0.1:9992 samplekey")
				}

				

			}else{

				fmt.Println("Looks like server is not running on "+host)
				fmt.Println("Try running server on "+host)
				fmt.Println("Example: ./cliclient get 127.0.0.1:9992 samplekey")
			}

		}else if action == "watch" {

			if verifyHost(host) == true {

				if(argslen == 4){

					key := os.Args[3]
					clientWatch(host, key)
					

				}else{

					fmt.Println("Please see correct usage of get command")
					fmt.Println("Example: ./cliclient watch 127.0.0.1:9992 samplekey")
				}

				

			}else{

				fmt.Println("Looks like server is not running on "+host)
				fmt.Println("Try running server on "+host)
				fmt.Println("Example: ./cliclient watch 127.0.0.1:9992 samplekey")
			}

		}else{

			fmt.Println(usage)

		}

	}





}