package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{
	
}

func RenderTemplate(w http.ResponseWriter, tmpl string){
	tc,err:=CreateTemplateCache()
	if err!=nil{
		log.Fatal(err)
	}
	t,ok:=tc[tmpl]
	if !ok{
		log.Fatal(err)
	}
	buf:=new(bytes.Buffer)
	_=t.Execute(buf,nil)
	_,err=buf.WriteTo(w)
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

