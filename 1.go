package custframe

import (
	"fmt"
	"github.com/liuhangkaixcode/custframe/tools"
)

func Hello()  {
	s:=tools.GetString()
	fmt.Println(s)
}

func V1()  {
	fmt.Println("===========V1")
}


func V2()  {
	fmt.Println("===========V2")
}
