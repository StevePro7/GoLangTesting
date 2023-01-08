package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strings"
)

func uppercasePodLabel(clientset kubernetes.Interface, namespace, podName, labelKey string) (string, error) {

	pod, err := clientset.CoreV1().Pods(namespace).Get(
		context.Background(),
		podName,
		v1.GetOptions{},
	)
	if err != nil {
		return "", err
	}

	labelValue, ok := pod.ObjectMeta.Labels[labelKey]
	if !ok {
		return "", fmt.Errorf("no label with key %s for pod %s/%s", labelKey, namespace, podName)
	}

	return strings.ToUpper(labelValue), nil
}

func main() {
	fmt.Println("unit testing resources")
}
