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

Useful for debugging:

```sh
$ dagger -c 'container | from alpine | with-secret-variable FOO env://GHCR_TOKEN | with-exec -- sh -c "echo secret is \$FOO" | stdout'
```

```sh
func (m *MyModule) SetEnvVar(ctx context.Context) (string, error) {
	return dag.Container().
		From("alpine").
		WithEnvVariable("ENV_VAR", "VALUE").    <------!
		WithExec([]string{"env"}).
		Stdout(ctx)
}
```
