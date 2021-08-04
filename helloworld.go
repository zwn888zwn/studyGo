package main

import (
	"fmt"
	"strings"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello returns a personalised greeting in a given language.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	//fmt.Println(Hello("world", ""))
	var tokenStr = "Bearer test123456"
	if strings.HasPrefix(tokenStr,"Bearer "){
		tokenStr = string(tokenStr[len("Bearer "):]) //在验证时去掉Bearer前缀
		fmt.Println(tokenStr)
	}else {
		fmt.Println("error")
	}

}