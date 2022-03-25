package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_httpHeaders_OnHttpResponseHeaders(t *testing.T) {
	stdErr, kill := startEnvoy(t, 8001)
	defer kill()
	req, err := http.NewRequest("GET", "http://localhost:18000/uuid", nil)
	require.NoError(t, err)
	require.Eventually(t, func() bool {
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return false
		}
		key := "hello"
		value := "kurtis"
		return res.Header.Get(key) == value
	}, 5*time.Second, time.Millisecond, stdErr.String())
}
