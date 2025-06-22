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
	"io"
	"net/http"
	"crypto/tls"
  "crypto/x509"

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

// Returns a list of pods from the Kubernetes API
func (m *SreAiAgent) GetPods(
	// The context
	ctx context.Context, 
	// A directory that contains the Kubernetes service account token and ca.crt
	// This is the directory that is automatically mounted in the Pod at /var/run/secrets/kubernetes.io/serviceaccount
	kubernetesServiceAccountDir *dagger.Directory,
	) (string, error) {

		body, err := m.GetKubeAPI(ctx, kubernetesServiceAccountDir,"/api/v1/pods")
    if err != nil {
        return "", fmt.Errorf("failed to call GetKubeAPI and read response body: %w", err)
    }

    return string(body), nil
}

// Returns a list of pods from the Kubernetes API
func (m *SreAiAgent) GetKubeAPI(
	// The context
	ctx context.Context, 
	// A directory that contains the Kubernetes service account token and ca.crt
	// This is the directory that is automatically mounted in the Pod at /var/run/secrets/kubernetes.io/serviceaccount
	kubernetesServiceAccountDir *dagger.Directory,
	// The Kubernetes API path, eg "/api/v1/pods" to get Pods 
	apiPath string,
	) (string, error) {
    // Read the service account token
    token, err := kubernetesServiceAccountDir.File("token").Contents(ctx)
    if err != nil {
        return "", fmt.Errorf("failed to read service account token: %w", err)
    }

    // Read the CA certificate and set up HTTP client
    caCertString, err := kubernetesServiceAccountDir.File("ca.crt").Contents(ctx)
    if err != nil {
        return "", fmt.Errorf("failed to read CA cert: %w", err)
    }
		caCert := []byte(caCertString)
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)
    tlsConfig := &tls.Config{RootCAs: caCertPool}
    client := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}}

    // Create and send the request
    req, err := http.NewRequest("GET", "https://kubernetes.default.svc" + apiPath, nil)
    if err != nil {
        return "", fmt.Errorf("failed to create request: %w", err)
    }
    req.Header.Set("Authorization", "Bearer "+string(token))

    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("failed to call Kubernetes API: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("failed to read response body: %w", err)
    }

    return string(body), nil
}
