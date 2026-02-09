package main

func main() {
	srv, err := InitServer()
	if err != nil {
		panic(err)
	}
	srv.Run()
}
