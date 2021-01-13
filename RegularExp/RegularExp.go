package main

import (
	"fmt"
	"strings"
	"regexp"
)

func mail(){
	var usrmail string
	fmt.Println("Enter Email address")
	fmt.Scanln(&usrmail)
	validate(usrmail)
}

func validate(s string){
	re:=regexp.MustCompile("^[a-zA-Z0-9]+@[a-zA-Z]+\\.[a-zA-Z]")
	fmt.Printf("%v",re.MatchString(s))
}
func main(){  
	var option string
	fmt.Println("Do you want validate Email(Y/N)")
	fmt.Scanln(&option)
	option=strings.ToUpper(option)
	if option!="Y"{
		return
	}
		mail()
}
