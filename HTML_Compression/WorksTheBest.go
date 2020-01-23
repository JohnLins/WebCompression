package main

import (
	"fmt"
  "net/http"
  "io/ioutil"
)

func main() {
/*fmt.Print(string(GetHtml("http://linsmedia.tk")))
fmt.Println("___________________________________")*/


	initial := GetHtml("http://golang.org")/*[]byte(`<!doctype html>
         <div>   <p>hh h</p>

	      <h1 class=\"hi\">hiii john</h1> 

	  <p>hi</p> `)*/

    //fmt.Print(string(initial))
	fmt.Println("_______________")
	for i := 0; i < len(initial); i++ {
		if initial[i] == 10 || initial[i] == 9{
			initial[i] = 0
		}
		
		if initial[i] == 32 && initial[i + 1] == 32{
			initial[i] = 0
		}
		fmt.Print(string(initial[i]))
	}
	
}


func GetHtml(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return html
}


/*
[]byte(`<!doctype html>       <div> 
	      <h1 class=\"hi\">hiii john</h1> 
	  <p>hi</p>  `)
*/
/*
var scriptTag bool
if initial[i:8] == "<script" {
  scriptTag = true
}

if initial[i:8] == "</script>" {
  scriptTag = false
}*/