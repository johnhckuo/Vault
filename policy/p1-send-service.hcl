# Dev servers have version 2 of KV secrets engine mounted by default, so will
# need these paths to grant permissions:

path "secret/data/*" {
  capabilities = ["deny"]
}


path "secret/data/dev/participant1/send-service/*" {
  capabilities = ["read", "update", "create"]
}

path "secret/data/dev/ww/pr-service/*" {
  capabilities = ["read", "update"]
}


