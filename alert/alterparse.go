package alert

import (
	"fmt"
	"reflect"
)

func AlterParse(alter map[string]int)(map[string]int,bool){
	fmt.Println("this is alter parse")
	fmt.Println(alter)
	fmt.Println(reflect.TypeOf(alter))
	var lasterdata map[string]int
	lasterdata = make(map[string]int)
	for k,v := range alter {
		if v > 0{
			lasterdata[k]=v
		}
	}
	if len(lasterdata) > 0{
		return lasterdata,true
	}
	return nil,false
}

