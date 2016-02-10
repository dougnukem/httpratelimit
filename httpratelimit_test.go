package httpratelimit_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dougnukem/httpratelimit"
)

var theAnswer = []byte("42")

func sleepHandler(timeout time.Duration) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(timeout)
			w.Write(theAnswer)
		})
}

func errorHandler(timeout time.Duration) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(timeout)
			w.WriteHeader(500)
			w.Write(theAnswer)
		})
}

func partialWriteNotifyingHandler(startedWriting chan<- struct{}, finishWriting <-chan struct{}) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte{theAnswer[0]})
			startedWriting <- struct{}{}
			<-finishWriting
			w.Write(theAnswer[1:])
		})
}

func assertResponse(res *http.Response, t *testing.T) {
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, theAnswer) {
		t.Fatalf(`did not find expected bytes "%s" instead found "%s"`, theAnswer, b)
	}
}

func TestOkWithDefaults(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(sleepHandler(time.Millisecond))
	defer server.Close()

	cachePolicy := httpratelimit.CacheByPath(time.Minute * 5)
	cache := httpratelimit.NewDefaultCache()

	transport := &httpratelimit.Transport{
		Config:    cachePolicy,
		ByteCache: cache,
		Transport: http.DefaultTransport,
	}
	client := &http.Client{Transport: transport}
	res, err := client.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	assertResponse(res, t)
}
