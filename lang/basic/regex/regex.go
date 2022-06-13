package main

import (
	"fmt"
	"regexp"
)

const text = `
my email is ccmouse@gmail.com@abc.com
email1 is abc@def.org
email2 is    kkk@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	regex := regexp.MustCompile(`([a-zA-Z\d]+)@([a-zA-Z\d]+)(\.[a-zA-Z\d.]+)`)
	//match := regex.FindString(text)
	//match := regex.FindAllString(text, -1)
	match := regex.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
