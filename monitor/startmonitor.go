package monitor

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"log"
	"k8s-client/httpclient"
	"fmt"
	"k8s-client/alert"
)

func StartMonitor(clientset *kubernetes.Clientset,client *http.Client){
	//以下为循环过程
	podNameList := pollPods(clientset)
	for _, podname := range podNameList {
		httpclient.GetPodRestartCountAarry(client, podname)
	}

	fmt.Println("PodRestartCountItems value is ", httpclient.PodRestartCount)

	//获取报警数据，如果取到数据，发送邮件，钉钉信息
	alertdata,b := alert.AlterParse(httpclient.PodRestartCount.PodRestartCountItems)
	if b{
		go alert.SendMail(alertdata)
		go alert.SendDingDing(alertdata)
	}
	//清空数据，否则会有脏数据
	httpclient.PodRestartCount.CleanUpPodRestartCountData()
}

//获取POD列表
func pollPods(clientset *kubernetes.Clientset) []string {
	//pod, err := clientset.CoreV1().Pods("default").Get("busybox", v1.GetOptions{})
	var podNameList []string
	pods, err := clientset.CoreV1().Pods("default").List(v1.ListOptions{})
	if err != nil {
		log.Println("fail to poll the pods and err is", err)
		return nil
	}

	for _, pod := range pods.Items {
		//fmt.Println(pod.Name)
		//podNameList.append(pod)
		podNameList = append(podNameList, pod.Name)
	}
	return podNameList
}