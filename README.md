# sre-ai-agent

And AI Agent that will act as an SRE and look after my cluster

## CI

To run the agent's CI

```sh
source .env

# Dagger Shell
publish ghcr.io tonygilkerson env://GHCR_TOKEN

# system shell
dagger -c 'publish ghcr.io tonygilkerson env://GHCR_TOKEN'

# Dagger Shell
dagger call publish --source=./agent --registry=ghcr.io --username=tonygilkerson --password=env://GHCR_TOKEN
```

## Secrets

Create the secrets for the dagger LLM

```sh
# be sure to source .env first
source .env

kubectl -n agent-wrk create secret generic dagger-llm \
  --from-literal=GEMINI_API_KEY="$GEMINI_API_KEY" \
  --from-literal=GEMINI_MODEL="$GEMINI_MODEL"
```

## Calling the Agent

```sh
# Dagger CLI
dagger call sre-ai-agent \
  --assignment="this is a test to see if you can call the getPods function. If you are successful at doing that then save a file to the archive folder named success.txt with the contents SUCCESS" \
  --archive-dir="." \
  --kubernetes-service-account-dir=../var/run/secrets/kubernetes.io/serviceaccount \
  export --path=.

dagger call sre-ai-agent \
  --assignment="Get a list of pods and save the pod list data in an archive file name pod-list.txt" \
  --archive-dir="." \
  --kubernetes-service-account-dir="." \
  export --path=.
 
dagger call sre-ai-agent \
  --assignment="Get a list of pods, extract relevant information such as pod name and namespace.  Save the extracted information in an archive file named pod-and-ns.json" \
  --archive-dir="." \
  --kubernetes-service-account-dir="." \
  export --path=.

dagger call sre-ai-agent \
  --assignment="This is a test to see if you can use the tools provided. First call GetPods to get a list of pods. Second, call ApplyJqFilter to extract the pod name and namespace. Third call WriteArchiveFile to save the extracted information in a file called pod-report.txt. Do NOT call  GetKubeAPI, ReadArchiveFile or ListArchiveFile. Those functions are not ready. " \
  --archive-dir="." \
  --kubernetes-service-account-dir="." \
  export --path=.

dagger call sre-ai-agent \
  --assignment="First call GetPods to get a list of pods. Second, call ApplyJqFilter to extract the pod name and namespace. Third use the extracted information to create a simple human readable report, Lastly call WriteArchiveFile to save the report in a file called pod-report.txt." \
  --archive-dir="." \
  --kubernetes-service-account-dir="." \
  export --path=.

dagger call sre-ai-agent \
  --assignment="First call GetPods to get a list of pods. Second, call ApplyJqFilter to extract the pod name and namespace in csv format. Third call RunAwkScript to format the extracted information into a simple human readable report. The report should have two columns with column heading, one for the pod name the second for the namespace. Lastly call WriteArchiveFile to save the report in a file called pod-report-hr.txt. Do not call readArchiveFile,  listArchiveFiles, or getKubeAPI they are not ready." \
  --archive-dir="." \
  --kubernetes-service-account-dir="." \
  export --path=.

dagger call sre-ai-agent \
  --assignment="Get a list of pods list of pods. You can use ApplyJqFilter to parse the kubernetes manifests in the pod list if needed. Inspect the pod manifests in the list and make recommendations for how they might better conform to best practices. Lastly call WriteArchiveFile to save the report in a file called pod-recommendations.txt. Do not call runAwkScript, readArchiveFile,  listArchiveFiles, or getKubeAPI they are not ready." \
  --archive-dir="." \
  --kubernetes-service-account-dir="." \
  export --path=.

# or dagger shell
resultsDir=$(sre-ai-agent "this is a test to see if you can call the getPods function. If you are successful at doing that then save a file to the archive folder named success.txt with the contents SUCCESS" "/" "../var/run/secrets/kubernetes.io/serviceaccount") 
container | from "busybox" | with-workdir "/wrk" | with-directory "/wrk" $resultsDir | with-exec cat ./archive/count.md | stdout


# dagger shell (testing)
# type "dagger" in the agent folder and then
sandbox "." "DEV"
```

## Useful for debugging:

```sh
$ dagger -c 'container | from alpine | with-secret-variable FOO env://GHCR_TOKEN | with-exec -- sh -c "echo secret is \$FOO" | stdout'
```../var/run/secrets/kubernetes.io/serviceaccount

```sh
func (m *MyModule) SetEnvVar(ctx context.Context) (string, error) {
	return dag.Container().
		From("alpine").
		WithEnvVariable("ENV_VAR", "VALUE").    <------!
		WithExec([]string{"env"}).
		Stdout(ctx)
}
```
