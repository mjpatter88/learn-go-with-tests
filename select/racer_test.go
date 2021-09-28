package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacing(t *testing.T) {
	t.Run("Returns the url of the fastest server", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer slowServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl)

		if got != want {
			t.Fatalf("got %q want %q", got, want)
		}
	})
	t.Run("Returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverA.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Fatalf("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
