package commonfunctions

import (
	"fmt"
	"io/ioutil"
	"os"
)

//ReadFile reads the file into a string
func ReadFile(filename string) string {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return string(bs)
}
