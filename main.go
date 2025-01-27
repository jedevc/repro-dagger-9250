package main

import (
	"dagger/example/internal/dagger"
)

type Example struct{}

func (m *Example) Test(mydir *dagger.Directory) *dagger.Container {
	return dag.Container().From("alpine").
		WithFile("entrypoint.sh", mydir.File("entrypoint.sh"), dagger.ContainerWithFileOpts{Permissions: 0500})
}
