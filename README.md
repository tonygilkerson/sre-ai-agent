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
  --assignment="Summarize the resource usage for all the pods in the cluster" \
  --archive-dir="/" \
  --kubernetes-service-account-dir=../var/run/secrets/kubernetes.io/serviceaccount

dagger call sre-ai-agent \
  --assignment="Get a list of pods in the cluster and use that list to count the number of pods.  Save the pod count in a simple markdown file called count.md then store that file in the archive folder" \
  --archive-dir="/" \
  --kubernetes-service-account-dir=../var/run/secrets/kubernetes.io/serviceaccount
# or dagger shell

resultsDir=$(sre-ai-agent "Get a list of pods in the cluster and use that list to count the number of pods.  Save the pod count in a simple markdown file called count.md then store that file in the archive folder" "/" "../var/run/secrets/kubernetes.io/serviceaccount") 
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
