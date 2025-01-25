### Just simple dockerized go app for k8s testing.

```
kubectl run demogo --image zavulon/devops-goslim:v1 --port=8080 --labels app=demogo
```

```
kubectl port-forward pod/demog 8080:8080

```
