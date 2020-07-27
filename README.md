Hashicorp Vault Basics

Assuming your vault is being hosted in k8s, and Kubernetes auth method of vault is enabled
```sh
kubectl port-forward -n vault vault-0 8200:8200
```

Then create a service account in k8s
```sh
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ServiceAccount
metadata:
  name: vault-user
  namespace: default
EOF
```

Check SA has been created successfully
```sh
kubectl get sa
```

Then you need to give your SA some policies to be able to read a KV path

```sh
vault write auth/kubernetes/role/demo \
    bound_service_account_names=vault-user \
    bound_service_account_namespaces=default \
    policies=<name-of-ur-policy-here> \
    ttl=1h
```

Retrieve the token of the service account
```sh
kubectl get secret <service-account-name>-<random digits> -o yaml
```

Set Envs
```sh
export VAULT_ADDR=https://127.0.0.1:8200
export TOKEN=<Token-We-Just-Created>
```

Start Running
```sh
go run main.go
```

