KUBERNETES_SERVICE_HOST and KUBERNETES_SERVICE_PORT. 



export SA_TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)

curl -v \
  --header "Authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
  --cacert /var/run/secrets/kubernetes.io/serviceaccount/ca.crt \
  https://kubernetes.default/api/v1/namespaces/agent-wrkroot/pods

curl -v   --header "Authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)"   --cacert /var/run/secrets/kubernetes.io/serviceaccount/ca.crt  https://kubernetes.default.svc/apis/metrics.k8s.io/v1beta1/pods
  
curl -v   --header "Authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)"   --cacert /var/run/secrets/kubernetes.io/serviceaccount/ca.crt  https://kubernetes.default.svc/apis/metrics.k8s.io/v1beta1/nodes
  

time curl -s -X POST http://192.168.50.196:11434/api/generate \
  -H "Content-Type: application/json" \
  -d '{
    "model": "qwen2.5-coder:14b",
    "prompt": "create a bash function to prompt the user for their name and print it to stdout"
  }' | jq -r '.response' | tr -d '\n'




env=$(env | with-container-input 'base' $base 'a base container' | with-container-output 'python-dev' 'a container with python dev tools')
llm | with-env $env | with-prompt "You have an alpine container. Install tools to develop with Python." | env | output python-dev | as-container | terminal