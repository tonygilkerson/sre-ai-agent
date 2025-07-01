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

// ReadArchiveFile reads the contents of an archive file located at the specified path within the archive directory.
// It takes a context for cancellation and timeout control, and the relative path to the archive file.
// Returns the file contents as a string, or an error if the file cannot be read.
func (m *Sandbox) ReadArchiveFile(
	// The context
	ctx context.Context,
	// The path to the archive file
	path string,
) (string, error) {

	return m.ArchiveDir.File(path).Contents(ctx)
}

// WriteArchiveFile adds a file with the specified name and content to the Sandbox's archive directory.
// It updates the ArchiveDir field by including the new file and returns the modified Sandbox instance.
//
// Parameters:
//   - ctx: The context for controlling cancellation and deadlines.
//   - fileName: The name of the file to be added to the archive.
//   - content: The contents of the file, represented as a *dagger.File.
//
// Returns:
//   - *Sandbox: The updated Sandbox instance with the new file included in the archive directory.
func (m *Sandbox) WriteArchiveFile(
	ctx context.Context,

	// The file name
	fileName string,

	// The contents of the file
	content *dagger.File,
) *Sandbox {

	m.ArchiveDir = m.ArchiveDir.WithFile(fileName, content)

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

// GetPods retrieves the list of Kubernetes pods as a Dagger file.
// It queries the Kubernetes API for the actual pod list and returns the result
// as a Dagger file named "pods.json".
// Panics if there is an error fetching the pod list from the Kubernetes API.
func (m *Sandbox) GetPods(ctx context.Context) *dagger.File {
	
	// This is a hack to make it easy to do some testing when run outside of a cluster
	_, err := m.KubernetesServiceAccountDir.File("token").Contents(ctx)
	if err != nil {
		println ("No token found, DEV mode assumed. Returning sample pod list")
		return dag.File("pods.json", string(samplePodList))
	}

	// Hit the k8s api for real
	body, err := m.GetKubeAPI(ctx, "/api/v1/pods")
	if err != nil {
		panic(err)
	}

	return dag.File("pods.json", string(body))

}

// GetKubeAPI retrieves data from the Kubernetes API server using the service account credentials
// mounted in the pod. It sends a GET request to the specified apiPath (e.g., "/api/v1/pods").
// The request is authenticated using the service account token and secured with the CA certificate
// provided by Kubernetes. The function returns the response body as a string, or an error if any
// step fails.
//
// Parameters:
//   - ctx: The context for controlling cancellation and timeouts.
//   - apiPath: The Kubernetes API path to query (e.g., "/api/v1/pods").
//
// Returns:
//   - string: The response body from the Kubernetes API server.
//   - error: An error if the request fails or if there are issues reading credentials or certificates.
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


// Run an awk script on the input content and return the results
// RunAwkScript executes an AWK script on the provided input file within a containerized environment.
//
// Parameters:
//   - ctx: The context for managing request lifetime and cancellation.
//   - script: The contents of the AWK script to be executed.
//   - content: The input file on which the AWK script will be applied.
//
// Returns:
//   - *dagger.File: The resulting file containing the output of the AWK script execution.
//   - error: An error if the execution fails.
//
// The function creates a container using the "busybox" image, writes the provided script and content
// to files within the container, and executes the AWK script against the content file. The output is
// captured and returned as a new file.
func (m *Sandbox) RunAwkScript(
	// The context
	ctx context.Context,

	// An awk script file. 
	script string,

	// The awk input file. The awk command will apply the `script` to this `content`
	content *dagger.File,

) (*dagger.File, error) {

	// Apply jq filter to content
	output, err := dag.Container().
		From("busybox").
		WithWorkdir("/wrk").
		WithFile("content.txt", content).
		WithNewFile("script.awk", script).
		WithExec([]string{"awk", "-f", "script.awk", "content.txt"}).
		Stdout(ctx)

	if err != nil {
    return nil, err
	}

	file := dag.Directory().WithNewFile("output.txt", output).File("output.txt")
	return file, nil
}


// RunKubeLinter runs kube-linter on the provided Kubernetes manifest file.
//
// It takes a context and a Dagger file containing Kubernetes manifest JSON data as input.
// The function installs kube-linter in a container, writes the manifest content to a file,
// and executes kube-linter to lint the manifest. The output from kube-linter is captured
// and returned as a Dagger file.
//
// Parameters:
//   - ctx: The context for controlling cancellation and deadlines.
//   - content: A Dagger file containing Kubernetes manifest JSON data.
//
// Returns:
//   - A Dagger file containing the kube-linter output.
//   - An error if any step in the process fails.
func (m *Sandbox) RunKubeLinter(
	// The context
	ctx context.Context,

	// A file containing Kubernetes manifest json data.
	content *dagger.File,

) (*dagger.File, error) {

	// Apply jq filter to content
	output, err := dag.Container().
		From("golang").
		WithWorkdir("/wrk").
		WithExec([]string{"go","install","golang.stackrox.io/kube-linter/cmd/kube-linter@latest"}).
		WithFile("content.json", content).
		WithExec([]string{"sh", "-c", "kube-linter lint content.json || true"}).
		Stdout(ctx)

	if err != nil {
    return nil, err
	}

	file := dag.Directory().WithNewFile("output.txt", output).File("output.txt")
	return file, nil
}