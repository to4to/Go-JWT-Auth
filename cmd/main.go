package main
import ("github.com/to4to/go-jwt-auth/cmd/api")

func main()  {
	server:=api.NewAPIServer("8080",nil)
	server.Run()
}