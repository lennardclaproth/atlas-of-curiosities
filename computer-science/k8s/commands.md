# Commands

## Context & Namespace

### Show current context

```bash
kubectl config current-context
```

### List contexts

```bash
kubectl config get-contexts
```

### Switch context

```bash
kubectl config use-context my-cluster
```

### List namespaces

```bash
kubectl get ns
```

### Set default namespace

```bash
kubectl config set-context --current --namespace=my-namespace
```

### Use namespace for a single command

```bash
kubectl get pods -n my-namespace
```

---

## Pods

### List pods

```bash
kubectl get pods
```

### List pods with node/IP information

```bash
kubectl get pods -o wide
```

### Watch pods

```bash
kubectl get pods -w
```

### Describe pod

```bash
kubectl describe pod my-pod
```

### Delete pod

```bash
kubectl delete pod my-pod
```

---

## Logs

### View logs

```bash
kubectl logs my-pod
```

### Follow logs

```bash
kubectl logs -f my-pod
```

### View previous container logs

```bash
kubectl logs my-pod --previous
```

### Logs from specific container

```bash
kubectl logs my-pod -c api
```

---

## Exec Into Containers

### Open shell

```bash
kubectl exec -it my-pod -- sh
```

or

```bash
kubectl exec -it my-pod -- bash
```

### Run command

```bash
kubectl exec my-pod -- ls -la
```

### Specific container

```bash
kubectl exec -it my-pod -c api -- sh
```

---

## Port Forwarding

### Forward a pod port

```bash
kubectl port-forward pod/my-pod 8080:80
```

### Forward a service port

```bash
kubectl port-forward svc/my-service 8080:80
```

Access:

```text
http://localhost:8080
```

### Multiple ports

```bash
kubectl port-forward svc/my-service 8080:80 8443:443
```

---

## Deployments

### List deployments

```bash
kubectl get deployments
```

### Describe deployment

```bash
kubectl describe deployment my-app
```

### Restart deployment

```bash
kubectl rollout restart deployment my-app
```

### Check rollout status

```bash
kubectl rollout status deployment my-app
```

### View rollout history

```bash
kubectl rollout history deployment my-app
```

### Roll back deployment

```bash
kubectl rollout undo deployment my-app
```

---

## Services

### List services

```bash
kubectl get svc
```

### Describe service

```bash
kubectl describe svc my-service
```

---

## Apply & Delete Resources

### Apply file

```bash
kubectl apply -f deployment.yaml
```

### Apply directory

```bash
kubectl apply -f k8s/
```

### Delete resource

```bash
kubectl delete -f deployment.yaml
```

### View rendered resource

```bash
kubectl get deployment my-app -o yaml
```

---

## Secrets & ConfigMaps

### List secrets

```bash
kubectl get secrets
```

### Describe secret

```bash
kubectl describe secret my-secret
```

### List configmaps

```bash
kubectl get configmaps
```

---

## Events

### Recent events

```bash
kubectl get events
```

### Sort by creation time

```bash
kubectl get events --sort-by=.metadata.creationTimestamp
```

---

## Resource Usage

### Node usage

```bash
kubectl top nodes
```

### Pod usage

```bash
kubectl top pods
```

### Namespace pod usage

```bash
kubectl top pods -n my-namespace
```

> Requires Metrics Server

---

## Storage

### List PVCs

```bash
kubectl get pvc
```

### Describe PVC

```bash
kubectl describe pvc my-pvc
```

### List PVs

```bash
kubectl get pv
```

### List storage classes

```bash
kubectl get storageclass
```

---

## Copy Files

### Pod → Local

```bash
kubectl cp my-pod:/tmp/file.txt ./file.txt
```

### Local → Pod

```bash
kubectl cp ./file.txt my-pod:/tmp/file.txt
```

---

## Debugging Workflow

### See everything in namespace

```bash
kubectl get all -n my-namespace
```

### Check pod status

```bash
kubectl get pods
```

Look for:

```text
CrashLoopBackOff
ImagePullBackOff
Error
Pending
```

### Investigate

```bash
kubectl describe pod my-pod
kubectl logs my-pod
kubectl logs -f my-pod
kubectl exec -it my-pod -- sh
```

---

## Useful Aliases

```bash
alias k=kubectl
```

Examples:

```bash
k get pods
k get svc
k logs -f my-pod
k exec -it my-pod -- sh
k port-forward svc/api 8080:80
```

---

## Daily Commands

```bash
kubectl config current-context

kubectl get pods -n my-namespace

kubectl describe pod <pod>

kubectl logs -f <pod>

kubectl exec -it <pod> -- sh

kubectl port-forward svc/<service> 8080:80

kubectl rollout restart deployment/<deployment>

kubectl get pvc

kubectl get storageclass
```
