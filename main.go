package fallfishunlimited

import (
	"io/ioutil"
	"net/http"
)

var (
	indexHtml []byte
)

func init() {
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		panic(err)
	}

	indexHtml = content

	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write(indexHtml)
}
