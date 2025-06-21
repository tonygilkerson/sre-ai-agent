// A generated module for SreAiAgentCi functions
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
	"dagger/sre-ai-agent-ci/internal/dagger"
	"fmt"
)

type SreAiAgentCi struct{}

// Build a ready-to-use development environment
func (m *SreAiAgentCi) BuildEnv(source *dagger.Directory) *dagger.Container {

	return dag.Container().
		From("ubuntu:latest").
		WithExec([]string{"sh", "-c", "apt-get update && apt-get install -y curl"}).
		WithExec([]string{"sh", "-c", "curl -fsSL https://dl.dagger.io/dagger/install.sh | BIN_DIR=/usr/local/bin sh"}).
		WithDirectory("/agent", source ).
		WithWorkdir("/agent")
}

// Publish the application container after building and testing it on-the-fly
func (m *SreAiAgentCi) Publish(
	ctx context.Context,
	// Registry address
	registry string,
	// Registry username
	username string,
	// Registry password
	password *dagger.Secret,
	// agent source
	source *dagger.Directory,
) (string, error) {

	return m.BuildEnv(source).
		WithRegistryAuth(registry, username, password).
		Publish(ctx, fmt.Sprintf("%s/%s/sre-ai-agent:latest", registry, username))
}


