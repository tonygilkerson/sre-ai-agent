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



`{"items":[{"metadata":{"name":"nginx-deployment-5d5b6b6b6b-abcde"}},{"metadata":{"name":"busybox-pod"}}]}`


 ! Error 429, Message: You exceeded your current quota, please check your plan and billing details. For more information on this error, head to:
    │   https://ai.google.dev/gemini-api/docs/rate-limits., Status: RESOURCE_EXHAUSTED, Details: [map[@type:type.googleapis.com/google.rpc.QuotaFailure
    │   violations:[map[quotaDimensions:map[location:global model:gemini-2.0-flash] quotaId:GenerateRequestsPerDayPerProjectPerModel-FreeTier
    │   quotaMetric:generativelanguage.googleapis.com/generate_content_free_tier_requests quotaValue:200]]] map[@type:type.googleapis.com/google.rpc.Help
    │   links:[map[description:Learn more about Gemini API quotas url:https://ai.google.dev/gemini-api/docs/rate-limits]]]
    │   map[@type:type.googleapis.com/google.rpc.RetryInfo retryDelay:4s]]


     ["sh", "-c", "jq -r '.items[].metadata.name' content.json > jqout"])co