package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func execFunc(input string) string {

	log.Println(string("exeFunc input : " + input))
	cmd := exec.Command("hellofunc")
	if cmd == nil {
		log.Println("cmd is nil")
	}

	stdin, err := cmd.StdinPipe()
	if nil != err {
		log.Printf("Error obtaining stdin: %s", err.Error())
	}

	stdout, err := cmd.StdoutPipe()
	if nil != err {
		log.Printf("Error obtaining stdout: %s", err.Error())
	}

	errStart := cmd.Start()
	if errStart != nil {
		log.Printf("cmd start error : %s", errStart.Error())
	}

	stdin.Write([]byte(input))
	stdin.Close()
	stdoutByte, _ := ioutil.ReadAll(stdout)
	cmd.Wait()

	log.Println("exeFunc output : " + string(stdoutByte))
	return string(stdoutByte)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {

	//inputStr := r.URL.Query()["msg"]
	inputStr, err := r.URL.Query()["msg"]

	if !err || len(inputStr) < 1 {
		log.Println("Url Param 'msg' is missing")
		return
	}

	log.Print(inputStr[0])

	//inputStr := r.FormValue("msg")

	retStr := execFunc(inputStr[0])

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(retStr))

	log.Print(retStr)
	log.Print([]byte(retStr))
	log.Println("Index page was called")
}

func handlerICon(w http.ResponseWriter, r *http.Request) {}

func main() {

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/favicon.ico", handlerICon)

	log.Println("Hello Cocktail Server start")
	http.ListenAndServe(":3000", nil)

}
