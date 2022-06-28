#!/bin/bash
echo "deleting...."
kubectl delete -f examples/backup.yaml

echo "building..."
go build /root/gopath/src/github.com/soda-cdm/kahu/cmd/kahu

echo "creating..."
kubectl create -f examples/backup.yaml