package e2e

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func SetupAPIHelper(t *testing.T) (close func()) {
	t.Helper()
	ctx := context.Background()
	container := prepareDBHelper(t, ctx)
	p, err := container.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.CommandContext(ctx, "go", "run", "../../cmd/api/main.go")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true, // Set process group ID to the same as the process ID
	}
	cmd.Env = append(cmd.Environ(), "HOST=localhost", fmt.Sprint("PORT=", p.Port()), "USER=postgres", "PASSWORD=postgres")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Start()

	require.NoError(t, err)

	waitForHealthCheck("http://localhost:1323/health", 1)

	return func() {
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		cmd.Wait()
		t.Log("close api server")
		t.Log("out:", stdout.String(), "err:", stderr.String())
	}

}

func waitForHealthCheck(url string, interval time.Duration) {

	for {
		// Make HTTP GET request to the health check API
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Error making request: %v\n", err)
		} else {
			// Check if status code is 200 OK
			if resp.StatusCode == http.StatusOK {
				log.Println("Health check passed with 200 OK")
				return
			}
			// Print the status code if not 200
			log.Printf("Received status code: %d\n", resp.StatusCode)
			resp.Body.Close()
			return
		}

		// Wait for the next interval before retrying
		time.Sleep(interval)
	}
}
