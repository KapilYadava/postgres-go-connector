# postgres-go-connector
This is a kubernetes application which connects postgres sql application running inside docker container outsider kubernetes cluster


# helpful links
https://stackoverflow.com/questions/43354167/minikube-expose-mysql-running-on-localhost-as-service/43477742#43477742

https://www.youtube.com/watch?v=fvpq4jqtuZ8

# Option-1

## headless-service 
## service without selectors
https://kubernetes.io/docs/concepts/services-networking/service/#services-without-selectors


# Option-2

## ExternalName service

apiVersion: v1
kind: Service
metadata:
  name: my-service
  namespace: prod
spec:
  type: ExternalName
  externalName: my.database.example.com

