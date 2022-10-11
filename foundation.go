package foundation

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewHandler() http.Handler {
	return mainHandler{}
}

type mainHandler struct {
}

func (m mainHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	webPage, _ := ioutil.ReadFile("./content/index.html")
	_, err := w.Write(webPage)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.WriteHeader(http.StatusOK)
}
