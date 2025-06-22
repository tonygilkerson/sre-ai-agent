KUBERNETES_SERVICE_HOST and KUBERNETES_SERVICE_PORT. 



export SA_TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)

curl -v \
  --header "Authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
  --cacert /var/run/secrets/kubernetes.io/serviceaccount/ca.crt \
  https://kubernetes.default/api/v1/namespaces/agent-wrkroot/pods

  curl -v   --header "Authorization: Bearer $(cat /var/run/secrets/kubernetes.io/serviceaccount/token)"   --cacert /var/run/secrets/kubernetes.io/serviceaccount/ca.crt  https://kubernetes.default.svc/pis/metrics.k8s.io/v1beta1/pods
  

  /var/run/secrets/kubernetes.io/serviceaccount