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
  --assignment="Get a list of pods in the cluster and use that list to create a pod report. This report should contain a human readable table that with one pod per row, summarizing the pod information. For each pod on the report include its name, the namespace it is contained in its resource allocation, eg. cpu and memory.  Save this report in markdown format in a file named pod-info-YYYY-MM-DD-HH-mm.md where YYYY-MM-DD-HH-mm is the actual Year, Month, Day, Hour and Minute when the report was created." \
  --archive-dir="." \
  --kubernetes-service-account-dir=../var/run/secrets/kubernetes.io/serviceaccount \
  export --path=.

# or dagger shell
resultsDir=$(sre-ai-agent "this is a test to see if you can call the getPods function. If you are successful at doing that then save a file to the archive folder named success.txt with the contents SUCCESS" "/" "../var/run/secrets/kubernetes.io/serviceaccount") 
container | from "busybox" | with-workdir "/wrk" | with-directory "/wrk" $resultsDir | with-exec cat ./archive/count.md | stdout
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
