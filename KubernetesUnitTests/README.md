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
https://stackoverflow.com/questions/57833499/how-to-write-simple-tests-for-client-go-using-a-fake-client
fake.New
k8s.io/client-go/kubernetes/fake
NewSimpleClientset

real client-go
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	
	config, err := rest.InClusterConfig()
	clientset, err := kubernetes.NewForConfig(config)
	ns, err := clientset.CoreV1().Namespaces().Get(namespace, metav1.GetOptions{})
    

IMPORTANT
check rs how I did it before because I may need to use KUBECONFIG
e.g. watch.go
cfgFile := os.Getenv("KUBECONFIG")
	cfg, err := clientcmd.BuildConfigFromFlags("", cfgFile)
	if err != nil {
		// attempt 2: in cluster config
		if cfg, err = rest.InClusterConfig(); err != nil {
			return false, err
		}
	}

	client, err := kubernetes.NewForConfig(cfg)



02.
https://medium.com/@e_frogers/unit-testing-with-kubernetes-client-go-283b11aaa7db
author uses this as a reference
https://github.com/kubernetes/ingress-nginx/blob/main/internal/k8s/main_test.go
	
_, err := k.Client.CoreV1().Namespaces().Create(ns)

similar code in
infra_k8s.go

k8s.io/client-go/kubernetes/fake
proxy_test.go
_, err = proxy.New(fake.NewSimpleClientset(), nil, "testnode")


03.
good article
Creating a Go Kubernetes client with client-go
https://dev.to/muaazsaleem/creating-a-go-kubernetes-client-with-client-go-jlf

marry up the IMPORTANT part above
i.e.
create out of cluster config
cfg, err := clientcmd.BuildConfigFromFlags("", cfgFile)
vs.
create in cluster config
cfg, err = rest.InClusterConfig()

manage package dependencies
k8s.io/api
k8s.io/apimachinery

AWESOME at the end
TODO series
Vladimir Vivien
https://medium.com/programming-kubernetes/building-stuff-with-the-kubernetes-api-toc-84d751876650
https://medium.com/programming-kubernetes/building-stuff-with-the-kubernetes-api-1-cc50a3642
https://medium.com/programming-kubernetes/building-stuff-with-the-kubernetes-api-part-4-using-go-b1d0e3c1c899


04.
Thomas Stringer
Unit Testing Kubernetes Resources with Go
https://trstringer.com/unit-test-kubernetes-resources-go

references this article as a pre-req
Connect to Kubernetes from Go
https://trstringer.com/connect-to-kubernetes-from-go

k8s.io/client-go/kubernetes/fake

