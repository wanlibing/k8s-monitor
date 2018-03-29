package logger

import(
"fmt"
"log"
"os"
)

func MLog()(logger *log.Logger){
	logfile,err:=os.OpenFile("k8smonitor.log",os.O_APPEND|os.O_CREATE,0666)
	if err!=nil{
		fmt.Printf("%s\r\n",err.Error())
		os.Exit(-1)
	}
	//defer logfile.Close()
	logger=log.New(logfile,"\r\n",log.Ldate|log.Ltime|log.Llongfile)
	return logger
}
