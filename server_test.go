package foundation

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandler_ServeHTTP(t *testing.T) {
	testHandler, responseRecorder, body := setupTestHTTPEnvironment()
	testServingHTTP(testHandler, responseRecorder, body)
	checkIfResponseCodeIsOk(t, responseRecorder)
}

func testServingHTTP(testHandler http.Handler, responseRecorder httptest.ResponseRecorder, body io.Reader) {
	testHandler.ServeHTTP(&responseRecorder, httptest.NewRequest(http.MethodGet, "localhost:8080", body))
}

func checkIfResponseCodeIsOk(t *testing.T, responseRecorder httptest.ResponseRecorder) {
	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("response code not 200: %d\n", responseRecorder.Code)
	}
}

func setupTestHTTPEnvironment() (http.Handler, httptest.ResponseRecorder, io.Reader) {
	var testHandler http.Handler = mainHandler{}
	responseRecorder := httptest.ResponseRecorder{}
	var body io.Reader
	return testHandler, responseRecorder, body
}
