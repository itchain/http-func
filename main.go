package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func execFunc(input string) string {

	cmd := exec.Command("./hellofunc")

	stdin, err := cmd.StdinPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdin: %s", err.Error())
	}

	stdout, err := cmd.StdoutPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdout: %s", err.Error())
	}

	cmd.Start()
	stdin.Write([]byte(input))
	stdin.Close()
	stdoutByte, _ := ioutil.ReadAll(stdout)
	cmd.Wait()

	//log.Println(string(stdoutByte))
	return string(stdoutByte)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {

	inputStr := r.URL.Query()["msg"][0]
	retStr := execFunc(inputStr)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(retStr))

	log.Print(retStr)
	log.Println("Index page was called")
}

func main() {

	http.HandleFunc("/", handleIndex)

	log.Println("Hello Cocktail Server start")
	fmt.Println(http.ListenAndServe(":3000", nil))

}
