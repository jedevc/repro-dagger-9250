package main

import (
	"dagger/example/internal/dagger"
)

type Example struct{}

func (m *Example) Test() *dagger.Container {
	src := dag.Directory().WithNewFile("entrypoint.sh", "#!/bin/sh\necho foo")
	return dag.Container().From("alpine").
		WithFile("entrypoint.sh", src.File("entrypoint.sh"), dagger.ContainerWithFileOpts{Permissions: 0500})
}
