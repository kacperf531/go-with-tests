package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const finnish = "Finnish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const finnishHelloPrefix = "Hei, "


func greetingPrefix(language string) (prefix string) {
    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
		prefix = spanishHelloPrefix
	case finnish:
		prefix = finnishHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
    return
}

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name
}

func main() {
    fmt.Println(Hello("world", ""))
}