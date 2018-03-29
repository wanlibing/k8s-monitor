package httpclient

import (
	"crypto/tls"
	"crypto/x509"
	//"log"
	//	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//	"reflect"
	"time"
)

var PodRestartCount *PodRestartCountData = &PodRestartCountData{
	Now: time.Now(),
	PodRestartCountItems: make(map[string]int),
}

func init(){
	PodRestartCount.SetInit()
}

func Getconn() *http.Client {
	pool := x509.NewCertPool()
	caCertPath := "/etc/kubernetes/ssl/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return nil
	}
	pool.AppendCertsFromPEM(caCrt)

	//cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key")
	cliCrt, err := tls.LoadX509KeyPair("/etc/kubernetes/ssl/client.crt", "/etc/kubernetes/ssl/client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return nil
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	return client
}

func GetPodRestartCountAarry(client *http.Client, podname string) {

	resp, err := client.Get("https://xx.xx.xxx.xx:6443/api/v1/namespaces/default/pods/" + podname)
	if err != nil {
		fmt.Println("http get error: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	PodRestartCount.AddPodRestartCountItems(string(body))
}
