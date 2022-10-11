package foundation

import "net/http"

type mainHandler struct {
}

func (m mainHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
