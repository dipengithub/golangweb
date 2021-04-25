package main
//go run .
//go mod init github.com/dip/djhcudn
import (
	handler "github.com/dipengithub/golangweb/pkg/handlers"
	"net/http"

)

func main(){
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	_=http.ListenAndServe(":8080",nil)
}

