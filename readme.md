### Just simple dockerized go app for k8s testing.

```
kubectl run demo-go --image zavulon/devops-goslim:v1 --port=8080 --labels app=demo-go
```

```
kubectl port-forward pod/demo-go 8080:8080

```
