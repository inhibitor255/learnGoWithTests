package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServerDelayTime := 20 * time.Millisecond
		fastServerDelayTime := 0 * time.Millisecond

		slowServer := makeDelayedServer(slowServerDelayTime)

		fastServer := makeDelayedServer(fastServerDelayTime)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}

	})
	t.Run("returns an error if a server doesn't respond within timeout", func(t *testing.T) {
		server := makeDelayedServer(15 * time.Millisecond)
		defer server.Close()
		var timeout = 10 * time.Millisecond

		_, err := ConfigurableRacer(server.URL, server.URL, timeout)

		if err == nil {
			t.Fatal("want error but none error get")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
