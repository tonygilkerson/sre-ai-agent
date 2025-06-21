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
	"fmt"
	 "os"
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
		WithExec([]string{"sh", "-c", "apt-get update && apt-get install -y curl"}).
		WithExec([]string{"sh", "-c", "curl -fsSL https://dl.dagger.io/dagger/install.sh | BIN_DIR=/usr/local/bin sh"}).
		WithWorkdir("/agent")
}

// Publish the application container after building and testing it on-the-fly
func (m *SreAiAgent) Publish(
	ctx context.Context,
	// Registry address
	registry string,
	// Registry username
	username string,
	// Registry password
	password *dagger.Secret,
) (string, error) {

	return m.BuildEnv().
		WithRegistryAuth(registry, username, password).
		Publish(ctx, fmt.Sprintf("%s/%s/sre-ai-agent:dev", registry, username))
}

// func (m *SreAiAgent) GetSaToken() (string,error) {
// 	data, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
// 	if err != nil {
// 			println("DEBUG err", err)
// 		return "", err
// 	}
// 	println("DEBUG SA token " + string(data))
// 	return string(data), nil
// } 