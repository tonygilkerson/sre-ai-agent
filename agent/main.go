// An SRE Agent use to look at the PODs in the cluster and make recommendations
package main

import (
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

	// The tools, inputs and outputs
	sandbox := dag.Sandbox(archiveDir, kubernetesServiceAccountDir)

	// The workspace for you to work on your assignment
	env := dag.Env().
		WithSandboxInput("sandbox", sandbox, "The tools you can use to complete your assignment.").
		WithStringInput("assignment", assignment, "The task to complete.").
		WithSandboxOutput("results", "The completed task results, eg. html and markdown reports.")

	// The agentic loop
	work := dag.LLM().
		WithEnv(env).
		WithPrompt(`
		You are an expert Kubernetes engineer on the Site Reliability Team.
		You have access to an archive directory where you can save your assignment results. 
		You can read past results from the archive directory if needed. This is good if you want to look up past assignment results. 
		Your assignment is : $assignment
		`)

	resultsDirectory := work.Env().Output("results").AsSandbox().ArchiveDir()

  // Return the Directory with the assignment completed
  return resultsDirectory
}
