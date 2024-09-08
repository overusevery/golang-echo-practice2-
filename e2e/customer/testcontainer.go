package e2e

import (
	"context"
	"testing"

	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

const (
	dbname   = "postgres"
	user     = "postgres"
	password = "postgres"
)

func prepareDB(t *testing.T) nat.Port {
	t.Helper()
	ctx := context.Background()

	// 1. Start the postgres container and run any migrations on it
	container, err := postgres.Run(
		ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase(dbname),
		postgres.WithUsername(user),
		postgres.WithPassword(password),
		postgres.BasicWaitStrategies(),
		postgres.WithSQLDriver("pq"),
		postgres.WithInitScripts("../../fixture/postgres/init/customer.sql"),
	)
	if err != nil {
		t.Fatal(err)
	}

	p, err := container.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatal(err)
	}

	// 2. Create a snapshot of the database to restore later
	// err = container.Snapshot(ctx, postgres.WithSnapshotName("test-snapshot"))
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// Clean up the container after the test is complete
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	return p
}
