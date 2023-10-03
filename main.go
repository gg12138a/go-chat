package main

// test: `nc 127.0.0.1 8888`
func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
