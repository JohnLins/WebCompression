package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

  //COMPRESS CSS
  initial := []byte("#div { height: 20vw; width: 40vw; }")
  for i := 0; i < len(initial); i++ {
    if initial[i] == 32 {
      initial[i] = 1
    }
    fmt.Print(string(initial[i]))
  }
  //fmt.Println(initial)

  fmt.Print("\n")

  compressHTML()
  

}

func compressHTML() {

  initial := []byte("<!doctype html> <div>  <h1 class=\"hi\">hiii john</h1>   <p>hi</p>   ")
  //GetHtml("http://golang.org")
  
  var lthan bool //<
  var gthan bool //>
  var lthanslash bool //</
  var gthanagain bool //>
  
  //var delete bool
  
  for i := 0; i < len(initial); i++ {

    //Get rid on indents
    if initial[i] == 10 {
      initial[i] = 1
    }
    //

    if initial[i] == 60 {
      lthan = true
      //i+=1
    }
    if initial[i] == 62 {
      gthan = true
    }
    if initial[i] == 60 && initial[i + 1] == 47 {
      lthanslash = true
          //i += 1
    }
    if gthan == true && initial[i] == 62 {
      gthanagain = true
    }
  
    if lthan == true && gthan == true && lthanslash == true && gthanagain == true && initial[i + 1] != 47 && initial[i] == 32 {
                    //CHANGE initial[i] TO "S"  
                   //remove(initial, i)
                 // fmt.Print("S")//append(initial[i], 77)//initial[i] = 77
          initial[i] = 1 //Replaces space with an M
      if initial[i] == 60 {
            lthan = false
            gthan = false
            lthanslash = false
            gthanagain = false
        }
    }
/*
    if i == len(initial) - 2 {
      break
    }
*/
  
      fmt.Print(string(initial[i]))
    }
  
                     /*   for j := 0; j < len(initial) + 1; j++ {
                        fmt.Print(string(initial[j]))
                        }*/
  }

/*
//COMPRESS CSS
compressCSS() {
  initial := "#div { height: 20vw;width: 40vw; }"
  for i := 0; i < len(initial); i++ {
    if initial[i] == 32 || initial[i] == 10 {
      fmt.Print("$")
    }
  }
}
*/

/*
func remove(s []byte, i int) []byte {
    s[i] = s[len(s)-1]
    // We do not need to put s[i] at the end, as it will be discarded anyway
    return s[:len(s)-1]
}*/

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

if initial[i] = 62 {
  for j := 0; ; j++ {
    if initial[i + j] == 60 {
      break
    }
  }
}
*/




/*
if "<" == true && ">" == true && initial[i] == 32 {
  initial[i] = 1
}
*/




/*
validatorError := errorcount(url)

compress()

if errorcount() == validatorError {
  
} */