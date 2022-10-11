package foundation

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandler_ServeHTTP_ShowHTML(t *testing.T) {
	testHandler, responseWriter, body := setupTestHTTPEnvironment()
	testServingHTTP(testHandler, &responseWriter, body)
	responseContents := responseWriter.Body.String()
	fileContentBytes, err := ioutil.ReadFile("./content/index.html")
	fileContents := fmt.Sprintf("%s", fileContentBytes)
	if err != nil {
		t.Fatal(err.Error())
	}
	if responseContents != fileContents {
		t.Fatalf("%s does not match %s", responseContents, fileContents)
	}

}
func TestMainHandler_ServeHTTP(t *testing.T) {
	testHandler, responseWriter, body := setupTestHTTPEnvironment()
	testServingHTTP(testHandler, &responseWriter, body)
	checkIfResponseCodeIsOk(t, responseWriter)
}
func TestNewHandler(t *testing.T) {
	_ = NewHandler().(mainHandler)
}

var _ http.Handler = mainHandler{}

func testServingHTTP(h http.Handler, w http.ResponseWriter, body io.Reader) {
	h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "localhost:1234", body))
}

func checkIfResponseCodeIsOk(t *testing.T, responseRecorder spyWriter) {
	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("response code not 200: %d\n", responseRecorder.Code)
	}
}

func setupTestHTTPEnvironment() (http.Handler, spyWriter, io.Reader) {
	var testHandler http.Handler = mainHandler{}
	var responseRecorder = spyWriter{}
	var body io.Reader
	return testHandler, responseRecorder, body
}

type body struct {
	text string
}

func (b body) String() string {
	return b.text
}

type spyWriter struct {
	Code int
	Body body
}

func (s *spyWriter) Header() http.Header {
	//TODO implement me
	panic("implement me")
}

func (s *spyWriter) Write(bytes []byte) (int, error) {
	s.Body.text = fmt.Sprintf("%s", bytes)
	return 0, nil
}

func (s *spyWriter) WriteHeader(statusCode int) {
	s.Code = statusCode
}
