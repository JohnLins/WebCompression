package main

import (
	"fmt"
  "net/http"
  "io/ioutil"
)

func compressHTML(url string) ([]byte, int) {
	var reduction int = 0; 
	remove := func(s []byte, i int) []byte {
		return append(s[:i], s[i+1:]...)
	}

	html := getHTML(url)
	for i := 0; i<len(html)-1; i++ {

    //return
    if html[i] == 13 {
      html = remove(html, i)
			reduction++;
    }

    //Space
		for html[i + 1] == 32 && html[i] == 32 {
			html = remove(html, i)
			reduction++;
		}
    

		}
		
		return html, reduction
}

func main() {

  var compressedHTML, reduction = compressHTML("http://linsintegrations.co")
	
	fmt.Println(string(compressedHTML))

  fmt.Println(reduction, "B")
	 
}

func getHTML(url string) []byte {
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
