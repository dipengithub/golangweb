package render

import (
	"bytes"
	"fmt"
	"github.com/dipengithub/golangweb/pkg/config"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{
	
}
var app *config.AppConfig

func NewTemplates(a *config.AppConfig)  {
	app=a

}

func RenderTemplate(w http.ResponseWriter, tmpl string){
	tc:=app.TemplateCache
	t,ok:=tc[tmpl]
	if !ok{
		log.Fatal("could not get template cache")
	}
	buf:=new(bytes.Buffer)
	_=t.Execute(buf,nil)
	_,err:=buf.WriteTo(w)
	if err!=nil{
		fmt.Println("error",err)
	}



	//for html template without template format in another file
	//parsedTemplate,_:=template.ParseFiles("./templates/"+tmpl)
	//err=parsedTemplate.Execute(w,nil)
	//if err!=nil {
	//	fmt.Println(err)
	//	return
	//}
}
func  CreateTemplateCache() (map[string]*template.Template,error){
	myCash:=map[string]*template.Template{}
	pages,err:=filepath.Glob("./templates/*.page.tmpl")
	if err!=nil{
		return myCash,err
	}
	for _,page := range pages {
		name:=filepath.Base(page)
	    //	create  template set
	    ts,err:=template.New(name).Funcs(functions).ParseFiles(page)
	    if err!=nil{
	    	return myCash,err
	       }
		matches,err:=filepath.Glob("./templates/*.layout.tmpl")
		if len(matches)>0{
			ts,err =ts.ParseGlob("./templates/*.layout.tmpl")
			if err!=nil{
				return myCash,err
			}
		}
		myCash[name]=ts

	}
	return myCash,nil


}

