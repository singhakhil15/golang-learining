package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error ")
		os.Exit(1)
	}

	fmt.Println(res.Status)
	fmt.Println(res.StatusCode)
	fmt.Println(res.Proto)

	// biteSlice, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(biteSlice))

	// bs := make([]byte, 999999)

	// res.Body.Read(bs)
	// fmt.Println(string(bs))

	// fmt.Println(" ------ ")
	// io.Copy(os.Stdout, res.Body)

	lw := logWriter{}

	io.Copy(lw, res.Body)

}

func (logWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}
