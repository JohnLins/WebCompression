package main

import (
	"fmt"
)

func remove(s []byte, i int) []byte {
    return append(s[:i], s[i+1:]...)
}

func main() {
html := []byte("<html><div>ddq eewat  twt<span>  </span><div>          </html>")

for i := 0; i<len(html)-1; i++ {
for html[i + 1] == 32 && html[i] == 32 {
	fmt.Println("Space")
	
	html = remove(html, i)
}
fmt.Println(string(html[i]))
}

fmt.Println(string(html))

}
