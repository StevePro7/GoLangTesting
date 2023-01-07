Kubernetes Unit Tests
06-Jan-2023

fake.NewSimpleClientset tutorial

https://www.sobyte.net/post/2022-04/k8s-unittest-guide
https://medium.com/@e_frogers/unit-testing-with-kubernetes-client-go-283b11aaa7db


Medium article but probably not so relevant
https://medium.com/the-phi/mocking-the-kubernetes-client-in-go-for-unit-testing-ddae65c4302


More info
Ooogle
fake.newsimpleclientset() hello world

https://pkg.go.dev/k8s.io/client-go/kubernetes/fake

https://stackoverflow.com/questions/57833499/how-to-write-simple-tests-for-client-go-using-a-fake-client
https://dev.to/muaazsaleem/creating-a-go-kubernetes-client-with-client-go-jlf
https://trstringer.com/unit-test-kubernetes-resources-go


Also interesting
Table driven tests
https://github.com/golang/go/wiki/TableDrivenTests


01.
Stack Overflow
fake.New
k8s.io/client-go/kubernetes/fake
NewSimpleClientset

real client-go
	client, err := kubernetes.NewForConfig(cfg)
	"k8s.io/client-go/kubernetes"

config, err := rest.InClusterConfig()




02.
https://medium.com/@e_frogers/unit-testing-with-kubernetes-client-go-283b11aaa7db
_, err := k.Client.CoreV1().Namespaces().Create(ns)

similar code in
infra_k8s.go

k8s.io/client-go/kubernetes/fake
proxy_test.go
_, err = proxy.New(fake.NewSimpleClientset(), nil, "testnode")
