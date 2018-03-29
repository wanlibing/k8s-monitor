package httpclient

import (
	"encoding/json"
	//"fmt"
	"sync"
	"time"
	"fmt"

	"reflect"

)

//var PodRestartCountItems []Container

type PodRestartCountData struct {
	mu sync.RWMutex
	PodRestartCountItems map[string]int
	Now time.Time
}

type Container struct {
	RestartCount int
	Name         string
}

type Status struct {
	ContainerStatuses []Container
}

type AAA struct {
	Status Status
}

func (p *PodRestartCountData) AddPodRestartCountItems(body string) {
	var s AAA
	str := body
	json.Unmarshal([]byte(str), &s)

	var status = s.Status
	c := &status.ContainerStatuses[0]
	p.PodRestartCountItems[c.Name]=c.RestartCount

}

func (p *PodRestartCountData) CleanUpPodRestartCountData(){
	p.mu.Lock()
	defer p.mu.Unlock()
	fmt.Println("delete PodRestartCountData")
	for k,_ := range p.PodRestartCountItems{
		delete(p.PodRestartCountItems,k)
	}

}

func (p *PodRestartCountData) ParseToJson() ( string, error) {
	fmt.Println("parse to json, and storge in elasesearch")
	data, err := json.Marshal(p.PodRestartCountItems)
	if err != nil{
		fmt.Println(err)
		return "",err
	}
	fmt.Println(data)
	fmt.Println(reflect.TypeOf(data))
	return string(data),nil
	}

func (p *PodRestartCountData) SetInit(){
	p = &PodRestartCountData{
		Now: time.Now(),
		PodRestartCountItems: nil,
	}

}
func (p *PodRestartCountData) SetTime(){
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Now = time.Now()
}