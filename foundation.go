package foundation

import (
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
	_, _ = w.Write(webPage)
	w.WriteHeader(http.StatusOK)
}
