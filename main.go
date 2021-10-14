package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type gostruct struct {
}

func main()  {
	f,err := os.OpenFile("./out.json",os.O_APPEND|os.O_CREATE|os.O_RDWR,0777)
	if err != nil{
		fmt.Println(err)
	}
	a := &gostruct{
	}
	ans,err := json.Marshal(a)
	if err != nil{
		fmt.Println(err)
	}
	f.Write(ans)
	f.Close()
	return
}