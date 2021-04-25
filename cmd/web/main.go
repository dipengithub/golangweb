package main
//go run .
//go mod init github.com/dip/djhcudn
import (
	"github.com/dipengithub/golangweb/pkg/config"
	handler "github.com/dipengithub/golangweb/pkg/handlers"
	"github.com/dipengithub/golangweb/pkg/render"
	"log"
	"net/http"

)

func main(){
	var app config.AppConfig
	tc,err:=render.CreateTemplateCache()
	if err!=nil{
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache=tc
	render.NewTemplates(&app)
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	_=http.ListenAndServe(":8080",nil)
}

