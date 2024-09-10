package e2e

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"syscall"
	"testing"

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

	//TODO//wait up server
	exec.Command("sleep", "10").Run()

	return func() {
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		cmd.Wait()
		t.Log("close api server")
		t.Log("out:", stdout.String(), "err:", stderr.String())
	}

}
