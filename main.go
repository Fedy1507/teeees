package main

func main() {
	msg := sayHello("Ali")

	println(msg)
}



func sayHello(name string) string {
	return "Hello " + name
}