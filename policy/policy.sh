vault policy write p1-send-policy p1-send-service.hcl
vault token create -policy=p1-send-policy -period=10h -explicit-max-ttl=20h -renewable=true
#vault token create -policy=p1-send-policy -period=15s -explicit-max-ttl=20h -renewable=true