package main

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s-client/httpclient"
	"net/http"
	"k8s-client/monitor"
	"time"
	"k8s-client/log"

)

var clientset *kubernetes.Clientset
var client *http.Client


func main() {
	kubeconfig := flag.String("kubeconfig", "./k8sconfig", "path to a kubeconfig. Requeire by out-of-cluster")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		logger.MLog().Fatalln(err)
	}
	//用k8s-client 与k8s连接
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {

		logger.MLog().Fatalln("something is wrong")
	}
	fmt.Println("##########################")
	//用http client 与k8s建立连接
	client = httpclient.Getconn()
	for {
		alertTimer := time.NewTimer(time.Second*5)
		if _,ok := <- alertTimer.C;ok{
			monitor.StartMonitor(clientset,client)
		}
		}

}



