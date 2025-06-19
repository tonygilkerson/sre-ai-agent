// A generated module for SreAiAgent functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/sre-ai-agent/internal/dagger"
)

type SreAiAgent struct{}

// Returns a container that echoes whatever string argument is provided
func (m *SreAiAgent) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *SreAiAgent) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

// Build a ready-to-use development environment
func (m *SreAiAgent) BuildEnv() *dagger.Container {

	return dag.Container().
		From("ubuntu:latest").
		WithWorkdir("/usr/local/bin").
		WithExec([]string{"sh", "-c", "apt-get update && apt-get install -y curl"}).
		WithExec([]string{"sh", "-c", "curl -fsSL https://dl.dagger.io/dagger/install.sh | BIN_DIR=/usr/local/bin sh"}).
		WithExec([]string{"sh", "-c", "curl -LO 'https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl'"}). 
		WithExec([]string{"chmod", "+x", "/usr/local/bin/kubectl"})

}


//export _EXPERIMENTAL_DAGGER_RUNNER_HOST="kube-pod://dagger-engine-smbm4?namespace=dagger&container=dagger-engine"