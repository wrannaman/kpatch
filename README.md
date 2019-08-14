# Kpatch 

## Before 

```sh
kubectl patch deployment api-deployment "{"spec":{"template":{"metadata":{"labels":{"creationTimestamp":"`date +'%s'`"}}}}}"
```

## After 

```sh
kpatch api-deployment
```

## Install 

```sh
go install https://github.com/wrannaman/kpatch 
```