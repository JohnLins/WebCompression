package main

import (
	"fmt"
)

func main() {

  initial := []byte("<div><h1>Hello World!</h1>   <h1>hu</h1>      <p>cvaafvar</p></div>")

var lthan bool //<
var gthan bool //>
var lthanslash bool //</
var gthanagain bool //>

//var delete bool

for i := 0; i < len(initial); i++ {
  if initial[i] == 60 {
    lthan = true
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
        initial[i] = 77 //Replaces space with an M
    if initial[i] == 60 {
        lthan = false
        gthan = false
        lthanslash = false
        gthanagain = false
        }
  }

    fmt.Print(string(initial[i]))
  }

}

/*
func remove(s []byte, i int) []byte {
    s[i] = s[len(s)-1]
    // We do not need to put s[i] at the end, as it will be discarded anyway
    return s[:len(s)-1]
}*/
