package service_test

import (
	"base-project/src/app"
	"context"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Chdir("../../../")

	app.Init(context.Background())

	exitVal := m.Run()

	os.Exit(exitVal)
}
