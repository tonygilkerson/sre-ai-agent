// An SRE Agent use to look at the PODs in the cluster and make recommendations
package main

import (
	"context"
	"dagger/sre-ai-agent/internal/dagger"
)

type SreAiAgent struct{}

// An agent use to inspect the workloads in a cluster and report on recommendations
func (m *SreAiAgent) SreAiAgent(
	// The task to perform
	assignment string,

	// The archive directory for storing results in persistent storage
	archiveDir *dagger.Directory,

	// A directory that contains the Kubernetes service account token and ca.crt
	// This is the directory that is automatically mounted in the Pod at /var/run/secrets/kubernetes.io/serviceaccount
	kubernetesServiceAccountDir *dagger.Directory,

) *dagger.Directory {

	// // The tools, inputs and outputs
	// sandbox := dag.Sandbox(archiveDir, kubernetesServiceAccountDir)

	// The workspace for you to work on your assignment
	env := dag.Env().
		WithSandboxInput("sandbox", dag.Sandbox(archiveDir, kubernetesServiceAccountDir), "The tools you can use to complete your assignment.").
		WithStringInput("assignment", assignment, "The task to complete.").
		WithSandboxOutput("results", "The completed task results, eg. html and markdown reports.")

	// The agentic loop
	work := dag.LLM().
		WithEnv(env).
		WithPrompt(`
		You are an expert Kubernetes engineer on the Site Reliability Team.
		You have access to an archive directory where you can save your assignment results. 
		Your assignment is : $assignment
		`)
		// WithPrompt(`
		//   This is a test see if you can call getPods. 
		// 	Inspect the json data received and count the number of pods found.
		// 	Save the count in the archive directory with the name count.txt"
		// `)
			//You can read past results from the archive directory if needed. This is good if you want to report on progress towards your recommendations. 

	resultsDirectory := work.Env().Output("results").AsSandbox().ArchiveDir()

  // Return the Directory with the assignment completed
  return resultsDirectory
}

// This is a test
func (m *SreAiAgent) Test1(
	ctx context.Context,

	// The archive directory for storing results in persistent storage
	archiveDir *dagger.Directory,

	// A directory that contains the Kubernetes service account token and ca.crt
	// This is the directory that is automatically mounted in the Pod at /var/run/secrets/kubernetes.io/serviceaccount
	kubernetesServiceAccountDir *dagger.Directory,

) (string, error) {

	sandbox := dag.Sandbox(archiveDir, kubernetesServiceAccountDir)

	sandbox = sandbox.WriteArchiveFile("test20250625.txt", "this is a test")

	
	listing, err := sandbox.ListArchiveFiles(ctx)

	if err != nil {
		return "", err
	}

	// Return the Directory with the assignment completed
	return listing, nil

}

// This is a test
func (m *SreAiAgent) Test2(
	ctx context.Context,

	// The archive directory for storing results in persistent storage
	archiveDir *dagger.Directory,

	// A directory that contains the Kubernetes service account token and ca.crt
	// This is the directory that is automatically mounted in the Pod at /var/run/secrets/kubernetes.io/serviceaccount
	kubernetesServiceAccountDir *dagger.Directory,

) *dagger.Container {

	sandbox := dag.Sandbox(archiveDir, kubernetesServiceAccountDir)

	sandbox.WriteArchiveFile("test20250625.txt", "this is a test")

	return dag.Container().
		From("alpine:3").
		WithDirectory("/wrk", archiveDir).
		WithWorkdir("/wrk")
}
// This is a test
func (m *SreAiAgent) Test3(
	ctx context.Context,

	// The archive directory for storing results in persistent storage
	archiveDir *dagger.Directory,

	// A directory that contains the Kubernetes service account token and ca.crt
	// This is the directory that is automatically mounted in the Pod at /var/run/secrets/kubernetes.io/serviceaccount
	kubernetesServiceAccountDir *dagger.Directory,

) (string, error) {

	sandbox := dag.Sandbox(archiveDir, kubernetesServiceAccountDir)

	// To "update" the archiveDir in the Sandbox, you need to create a new Sandbox with the updated Directory,
	// since you cannot assign to a field directly if it's not exported or is not addressable.
	updatedArchiveDir := sandbox.ArchiveDir().WithNewFile("test20250625.txt", "this is a test")

	// If you want the Sandbox to use the new archiveDir, you must construct a new Sandbox:
	updatedSandbox := dag.Sandbox(updatedArchiveDir, kubernetesServiceAccountDir)

	return dag.Container().
		From("alpine:3").
		WithDirectory("/src", updatedSandbox.ArchiveDir()).
		WithWorkdir("/src").
		WithExec([]string{"tree", "."}).
		Stdout(ctx)
}
// This is a test
func (m *SreAiAgent) Test4(
	ctx context.Context,

	// The archive directory for storing results in persistent storage
	archiveDir *dagger.Directory,

	// A directory that contains the Kubernetes service account token and ca.crt
	// This is the directory that is automatically mounted in the Pod at /var/run/secrets/kubernetes.io/serviceaccount
	kubernetesServiceAccountDir *dagger.Directory,

) (string, error) {

	sandbox := dag.Sandbox(archiveDir, kubernetesServiceAccountDir)

	// To "update" the archiveDir in the Sandbox, you need to create a new Sandbox with the updated Directory,
	// since you cannot assign to a field directly if it's not exported or is not addressable.
	updatedArchiveDir := sandbox.ArchiveDir().WithNewFile("test20250625.txt", "this is a test")

	return dag.Container().
		From("alpine:3").
		WithDirectory("/src", updatedArchiveDir).
		WithWorkdir("/src").
		WithExec([]string{"tree", "."}).
		Stdout(ctx)
}
