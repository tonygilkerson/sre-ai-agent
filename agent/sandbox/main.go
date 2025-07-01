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
	}

	file := dag.Directory().WithNewFile("output.txt", output).File("output.txt")
	return file, nil
}

// Run an awk script on the input content and return the results
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

// DEVTODO stackrox/kube-linter:latest
