package custframe

import (
	"fmt"

	"github.com/liuhangkaixcode/custframe/tools"
)

func Hello() {
	s := tools.GetString()
	fmt.Println(s)
}

func Master() {
	fmt.Println("===========Master")
}
