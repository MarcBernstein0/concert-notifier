package spotify

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var server *httptest.Server

func TestMain(m *testing.M) {
	fmt.Println("Mock Server")
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		trimPath := strings.TrimSpace(r.URL.Path)
		switch trimPath {
		default:
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	}))

	fmt.Println("run tests")
	m.Run()
}

func TestCreateCustomClient(t *testing.T) {
	// Given
	givenCustomClient := customClient{
		baseURL: "testEndpoint",
		client:  http.DefaultClient,
		config: struct {
			clientId     string
			clientSecret string
		}{
			clientId:     "testId",
			clientSecret: "testSecret",
		},
		timeOut: 20,
	}
	// When
	res := New("testEndpoint", "testId", "testSecret", http.DefaultClient, 20)
	// Then
	require.Equal(t, givenCustomClient, res)
}

func TestFetchSpotifyTopArtist(t *testing.T) {
	// Given
	res := New(server.URL, "testId", "testSecret", http.DefaultClient, 20)
	expectedRes := []string{
		"artist1",
		"artist2",
		"artist3",
	}

	// When
	givenRes, givenErr := res.Fetch(context.Background())

	// Then
	require.NoError(t, givenErr)
	require.ElementsMatch(t, expectedRes, givenRes)
}
