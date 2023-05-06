package main

import (
	"code-generator-test/pkg/generated/clientset/versioned"
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	list, err := clientset.HjjzsV1().Foos("default").Get(context.TODO(), "hjj", v1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", list)
}
