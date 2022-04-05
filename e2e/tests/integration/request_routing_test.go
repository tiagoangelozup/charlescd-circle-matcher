//go:build integration
// +build integration

package integration

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_request_routing_1(t *testing.T) {
	ns := os.Getenv("NAMESPACE")
	require.NotEmpty(t, ns)
	url := fmt.Sprintf("http://%s.lvh.me/", ns)
	log.Printf("trying to get 'RED' page on %q\n", url)
	req, err := http.NewRequest("GET", url, nil)
	require.NoError(t, err)

	require.Eventually(t, getRed(req), 5*time.Second, 500*time.Millisecond)
	require.Never(t, getBlue(req), 5*time.Second, 500*time.Millisecond)
}

func Test_request_routing_2(t *testing.T) {
	ns := os.Getenv("NAMESPACE")
	require.NotEmpty(t, ns)

	url := fmt.Sprintf("http://%s.lvh.me/", ns)
	log.Printf("trying to get 'BLUE' page on %q\n", url)
	req, err := http.NewRequest("GET", url, nil)
	require.NoError(t, err)

	user := base64.RawStdEncoding.EncodeToString([]byte(`{"name":"Rafaela Rocha Cavalcanti","age":33,"city":"Fortaleza-CE"}`))
	req.Header.Add("X-CharlesCD-User", user)
	require.Eventually(t, getBlue(req), 5*time.Second, 500*time.Millisecond)
	require.Never(t, getRed(req), 5*time.Second, 500*time.Millisecond)
}

func Test_request_routing_3(t *testing.T) {
	ns := os.Getenv("NAMESPACE")
	require.NotEmpty(t, ns)

	url := fmt.Sprintf("http://%s.lvh.me/", ns)
	log.Printf("trying to get 'RED' page on %q\n", url)
	req, err := http.NewRequest("GET", url, nil)
	require.NoError(t, err)

	user := base64.RawStdEncoding.EncodeToString([]byte(`{"name":"Antônio Rodrigues Santos","age":18,"city":"Lençóis Paulista-SP"}`))
	req.Header.Add("X-CharlesCD-User", user)
	require.Eventually(t, getRed(req), 5*time.Second, 500*time.Millisecond)
	require.Never(t, getBlue(req), 5*time.Second, 500*time.Millisecond)
}

func getRed(req *http.Request) func() bool {
	return func() bool {
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return false
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return false
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return false
		}
		return strings.Contains(string(body), "background-color:red;")
	}
}

func getBlue(req *http.Request) func() bool {
	return func() bool {
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return false
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return false
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return false
		}
		return strings.Contains(string(body), "background-color:blue;")
	}
}
