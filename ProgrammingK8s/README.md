Building stuff with the Kubernetes API (Part 4) — Using Go
21-Sep-2023

https://medium.com/programming-kubernetes/building-stuff-with-the-kubernetes-api-part-4-using-go-b1d0e3c1c899


https://github.com/vladimirvivien/k8s-client-examples/blob/master/go/pvcwatch/main.go
https://github.com/vladimirvivien/k8s-client-examples/blob/master/go/pvcwatch-ctl/main.go

https://github.com/vladimirvivien/k8s-client-examples/blob/master/go/podlist/maing.go
module github.com/stevepro/k8s-client-example


podlist
go get golang.org/x/time
go get k8s.io/api 
go get k8s.io/apimachinery 
go get k8s.io/client-go
go get k8s.io/klog
go get k8s.io/utils



pvcwatch
go get golang.org/x/time
go get k8s.io/api 
go get k8s.io/apimachinery 
go get k8s.io/client-go
go get k8s.io/klog
go get k8s.io/utils




pvcwatch-cli
