//go:build integration
// +build integration

package integration

import (
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
	require.Eventually(t, func() bool {
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
	}, 5*time.Second, time.Millisecond)
}

func Test_request_routing_2(t *testing.T) {
	ns := os.Getenv("NAMESPACE")
	require.NotEmpty(t, ns)

	url := fmt.Sprintf("http://%s.lvh.me/", ns)
	log.Printf("trying to get 'BLUE' page on %q\n", url)
	req, err := http.NewRequest("GET", url, nil)
	require.NoError(t, err)
	req.Header.Add("end-user", "18e2451b-c1c3-4d4f-a37b-232db6e95cf9")
	require.Eventually(t, func() bool {
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
	}, 5*time.Second, time.Millisecond)
}
