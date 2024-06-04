package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ReadJson(filename string) (map[string]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var info map[string]interface{}
	err = json.Unmarshal(content, &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var name string
		switch r.Method {
		case "POST":
			name = r.PostFormValue("name")
			fmt.Println("POST name:", name)
		case "GET":
			name = r.URL.Query().Get("name")
			fmt.Println("GET name:", name)
		default:
			http.Error(w, "접근 불가 메서드", http.StatusMethodNotAllowed)
			return
		}
		fmt.Println(name)
		professormap, err := ReadJson("professorinfo.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)
			return
		}
		value, ishas := professormap[name]
		if !ishas {
			http.Error(w, fmt.Sprintf("%s 교수님을 찾을수 없습니다.", name), http.StatusNotFound)
			return
		}
		fmt.Fprint(w, value)

	})
	http.ListenAndServe("0.0.0.0:8000", nil)
}
