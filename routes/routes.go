package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Response struct {
	LLMissRate  string `json:"LL miss rate"`
	Branches    string `json:"Branches"`
	I1Misses    string `json:"I1  misses"`
	I1MissRate  string `json:"I1  miss rate"`
	LLiMissRate string `json:"LLi miss rate"`
	DRefs       string `json:"D   refs"`
	LLdMisses   string `json:"LLd misses"`
	LLMisses    string `json:"LL misses"`
	Mispredicts string `json:"Mispredicts"`
	MispredRate string `json:"Mispred rate"`
	IRefs       string `json:"I   refs"`
	D1Misses    string `json:"D1  misses"`
	D1MissRate  string `json:"D1  miss rate"`
	LLdMissRate string `json:"LLd miss rate"`
	LLRefs      string `json:"LL refs"`
	LLiMisses   string `json:"LLi misses"`
}

func GetCFileAndPerform(w http.ResponseWriter, r *http.Request) {
	file, _, _ := r.FormFile("fileupload")
	defer file.Close()

	f, _ := os.OpenFile("./demo.c", os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	io.Copy(f, file)

	CompileACProgram("./demo.c")
	RunValgrind("demo")

	finalVal := GetValuesFromCallgrindFile("./demo.out")
	jsonString, _ := json.Marshal(finalVal)

	var result *Response
	json.Unmarshal(jsonString, &result)

	RespondWithJson(w, http.StatusOK, result)
}
