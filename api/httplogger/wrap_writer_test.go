package httplogger

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptest"
	"path"
	"runtime"
	"testing"
	"time"

	"golang.org/x/net/http2"
)

// NOTE: we must import `golang.org/x/net/http2` in order to explicitly enable
// http2 transports for certain tests. The runtime pkg does not have this dependency
// though as the transport configuration happens under the hood on go 1.7+.

var testdataDir string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	testdataDir = path.Join(path.Dir(filename), "../../testdata")
}

func TestWrapWriterHTTP2(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, fl := w.(http.Flusher)
		if !fl {
			t.Fatal("request should have been a http.Flusher")
		}
		_, hj := w.(http.Hijacker)
		if hj {
			t.Fatal("request should not have been a http.Hijacker")
		}
		_, rf := w.(io.ReaderFrom)
		if rf {
			t.Fatal("request should not have been a io.ReaderFrom")
		}
		_, ps := w.(http.Pusher)
		if !ps {
			t.Fatal("request should have been a http.Pusher")
		}

		w.Write([]byte("OK"))
	})

	wmw := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(NewWrapResponseWriter(w, r.ProtoMajor), r)
		})
	}

	server := http.Server{
		Addr:    ":7072",
		Handler: wmw(handler),
	}
	// By serving over TLS, we get HTTP2 requests
	go server.ListenAndServeTLS(testdataDir+"/cert.pem", testdataDir+"/key.pem")
	defer server.Close()
	// We need the server to start before making the request
	time.Sleep(100 * time.Millisecond)

	client := &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				// The certificates we are using are self signed
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Get("https://localhost:7072")
	if err != nil {
		t.Fatalf("could not get server: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("non 200 response: %v", resp.StatusCode)
	}
}

func TestFlushWriterRemembersWroteHeaderWhenFlushed(t *testing.T) {
	f := &flushWriter{basicWriter{ResponseWriter: httptest.NewRecorder()}}
	f.Flush()

	if !f.wroteHeader {
		t.Fatal("want Flush to have set wroteHeader=true")
	}
}

func TestHttpFancyWriterRemembersWroteHeaderWhenFlushed(t *testing.T) {
	f := &httpFancyWriter{basicWriter{ResponseWriter: httptest.NewRecorder()}}
	f.Flush()

	if !f.wroteHeader {
		t.Fatal("want Flush to have set wroteHeader=true")
	}
}

func TestHttp2FancyWriterRemembersWroteHeaderWhenFlushed(t *testing.T) {
	f := &http2FancyWriter{basicWriter{ResponseWriter: httptest.NewRecorder()}}
	f.Flush()

	if !f.wroteHeader {
		t.Fatal("want Flush to have set wroteHeader=true")
	}
}
