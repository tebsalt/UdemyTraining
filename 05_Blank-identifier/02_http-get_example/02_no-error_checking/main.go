package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.mcleods.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println("$s", page)
}

/* the blank identifier in this case, says i know an error
is coming back, but i'm not gonna do anything about it, so
just throw it away, hence why "err" is blank
*/
