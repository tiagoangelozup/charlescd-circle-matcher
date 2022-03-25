package e2e_test

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func startEnvoy(t *testing.T, adminPort int, cfg string) (stdErr *bytes.Buffer, kill func()) {
	cmd := exec.Command("envoy",
		"--base-id", strconv.Itoa(adminPort),
		"--concurrency", "1",
		"--component-log-level", "wasm:trace",
		"-c", cfg)

	buf := new(bytes.Buffer)
	cmd.Stderr = buf
	require.NoError(t, cmd.Start())
	require.Eventually(t, func() bool {
		res, err := http.Get(fmt.Sprintf("http://localhost:%d/listeners", adminPort))
		if err != nil {
			return false
		}
		defer res.Body.Close()
		return res.StatusCode == http.StatusOK
	}, 5*time.Second, 100*time.Millisecond, "Envoy has not started: "+stdErr.String())
	return buf, func() { require.NoError(t, cmd.Process.Kill()) }
}
