// A sandbox for the AI Agent
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"dagger/sandbox/internal/dagger"
	"fmt"
	"io"
	"net/http"
)

type Sandbox struct {
	// Container for publishing results
	Container *dagger.Container

	// Archive director for saving results for later retrieval
	Archive *dagger.Directory

	// A directory that contains the Kubernetes service account token and ca.crt
	// This is the directory that is automatically mounted in the Pod at /var/run/secrets/kubernetes.io/serviceaccount
	KubernetesServiceAccountDir *dagger.Directory
}

func New(archiveDir *dagger.Directory, kubernetesServiceAccountDir *dagger.Directory) *Sandbox {
	// A container to store the results
	results := dag.Container().From("nginx:latest").WithWorkdir("/usr/share/nginx/html")

	return &Sandbox{
		Container:                   results,
		Archive:                     archiveDir,
		KubernetesServiceAccountDir: kubernetesServiceAccountDir,
	}
}

// Read a file from the archive directory
func (m *Sandbox) ReadArchiveFile(
	// The context
	ctx context.Context,
	// The path to the archive file
	path string,
) (string, error) {
	
	return m.Archive.File(path).Contents(ctx)
}

// Write a file to the archive directory
func (m *Sandbox) WriteArchiveFile(
	// The path to the archive file
	path string,
	// The contents of the file
	content string,
) *Sandbox {

	m.Archive.WithNewFile(path, content)
	return m
}

// Read a file from the container
func (m *Sandbox) ReadContainerFile(
	// The context
	ctx context.Context,
	// The path to the file we want to read
	path string,
) (string, error) {

	return m.Container.File(path).Contents(ctx)
}

// Write a file to the container
func (m *Sandbox) WriteContainerFile(
	// The path to the file we want to write
	path string,
	// The contents of the file
	content string,
) *Sandbox {

	m.Container.WithNewFile(path, content)
	return m
}

// Returns a list of pods from the Kubernetes API
func (m *Sandbox) GetPods(ctx context.Context) (string, error) {

	body, err := m.GetKubeAPI(ctx, "/api/v1/pods")
	if err != nil {
		return "", fmt.Errorf("failed to call GetKubeAPI and read response body: %w", err)
	}

	return string(body), nil
}

// Returns a list of pods from the Kubernetes API
func (m *Sandbox) GetKubeAPI(
	// The context
	ctx context.Context,
	// The Kubernetes API path, eg "/api/v1/pods" to get Pods
	apiPath string,
) (string, error) {

	// Read the service account token
	token, err := m.KubernetesServiceAccountDir.File("token").Contents(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to read service account token: %w", err)
	}

	// Read the CA certificate and set up HTTP client
	caCertString, err := m.KubernetesServiceAccountDir.File("ca.crt").Contents(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to read CA cert: %w", err)
	}
	caCert := []byte(caCertString)
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{RootCAs: caCertPool}
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}}

	// Create and send the request
	req, err := http.NewRequest("GET", "https://kubernetes.default.svc"+apiPath, nil)
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
