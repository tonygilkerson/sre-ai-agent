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

	// Archive director for saving results for later retrieval
	ArchiveDir *dagger.Directory

	// A directory that contains the Kubernetes service account token and ca.crt
	// This is the directory that is automatically mounted in the Pod at /var/run/secrets/kubernetes.io/serviceaccount
	KubernetesServiceAccountDir *dagger.Directory

}

func New(archiveDir *dagger.Directory, kubernetesServiceAccountDir *dagger.Directory) *Sandbox {
	return &Sandbox{
		ArchiveDir:                  archiveDir,
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

	return m.ArchiveDir.File(path).Contents(ctx)
}

// Write a file to the archive directory
func (m *Sandbox) WriteArchiveFile(
	ctx context.Context,

	// The contents of the file
	content *dagger.Directory,
) *Sandbox {

	m.ArchiveDir = content

	return m
}

// List all of the files in the Archive
func (m *Sandbox) ListArchiveFiles(ctx context.Context) (string,error) {
	return dag.Container().
		From("alpine:3").
		WithDirectory("/src", m.ArchiveDir).
		WithWorkdir("/src").
		WithExec([]string{"ls", "-la", "."}).
		Stdout(ctx)

}

// Returns a list of pods from the Kubernetes API
func (m *Sandbox) GetPods(ctx context.Context) *dagger.Directory {
	
	// This is a hack to make it easy to do some testing when run outside of a cluster
	_, err := m.KubernetesServiceAccountDir.File("token").Contents(ctx)
	if err != nil {
		println ("No token found, DEV mode assumed. Returning sample pod list")
		return dag.Directory().WithNewFile("pods.json", string(samplePodList))
	}

	// Hit the k8s api for real
	body, err := m.GetKubeAPI(ctx, "/api/v1/pods")
	if err != nil {
		panic(err)
	}

	return dag.Directory().WithNewFile("pods.json", string(body))

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

// Apply a jq filter to on the input content and return the results
func (m *Sandbox) ApplyJqFilter(
	// The context
	ctx context.Context,

	// A jq filter expression, some examples: '.', '.name', '.[]', '.[] | select(.status == "active")', etc... 
	jqFilter string,

	// The JSON contents to apply a filter to
	content *dagger.Directory,

) *dagger.Directory {

	fmt.Printf("DEBUG jqFilter %s, content %s", jqFilter, content)

	// Apply jq filter to content
	return dag.Container().
		From("stedolan/jq").
		WithWorkdir("/wrk").
		WithDirectory("/wrk",content).
		WithExec([]string{"jq", jqFilter, "pods.json", ">", "filter-results.json"}).
		Directory("/wrk")
}

// validate json content
func (m *Sandbox) ValidateJson(
	// The context
	ctx context.Context,

	// The JSON contents to apply a filter to
	content *dagger.Directory,
) bool {

	fmt.Printf("DEBUG got content: %s", content)

	if ctx == nil{
		fmt.Println("no CTX")
	}

	// Apply jq filter to content
	return true

}