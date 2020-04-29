package routes

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func GetValuesFromCallgrindFile(filename string) map[string]string {
	readFile, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	raw_string := ""
	for _, eachline := range fileTextLines {
		raw_string += eachline + "\n"
	}

	pid_index := strings.Index(raw_string, "Cachegrind, a cache and branch-prediction profiler")

	if pid_index == -1 {
		log.Fatalf("Error occured while processing")
	}

	pid := raw_string[2 : pid_index-3]

	ind := strings.Index(raw_string, "I   refs:")
	new_string := strings.ReplaceAll(raw_string[ind:], "=="+pid+"== ", "")

	value_regex, _ := regexp.Compile(":(.*)")

	values := value_regex.FindAllString(new_string, -1)

	var value_slice []string
	for _, val := range values {
		whitespace := strings.ReplaceAll(val, " ", "")
		comma := strings.ReplaceAll(whitespace, ",", "")
		value_slice = append(value_slice, comma[1:])
	}

	key_regex, _ := regexp.Compile("(.*):")

	keys := key_regex.FindAllString(new_string, -1)

	var key_slice []string
	for _, k := range keys {
		key_slice = append(key_slice, k[:len(k)-1])
	}

	finalMap := make(map[string]string)
	for i := 0; i < len(key_slice); i += 1 {
		finalMap[key_slice[i]] = value_slice[i]
		// fmt.Println(key_slice[i], " = ", value_slice[i])
	}

	return finalMap
}

func CompileACProgram(filename string) {
	cmd := exec.Command("gcc", filename, "-o", filename[:len(filename)-2])
	stdout, _ := cmd.Output()

	if string(stdout) == "" {
		return
	}
}

func RunValgrind(filename string) {
	cmd := exec.Command("valgrind", "--tool=cachegrind", "--log-file="+filename+".out", "--branch-sim=yes", "./"+filename)
	stdout, _ := cmd.Output()

	fmt.Println("Output of the program -> ")
	fmt.Print(string(stdout) + "\n \n")
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
