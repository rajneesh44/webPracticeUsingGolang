package main 
import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var dir= "/home/rajneesh/go/src/web/basic/1/*"

func init(){
	tpl = template.Must(template.ParseGlob(dir))
}

func main(){
	// tpl, err := template.ParseFiles("tpl.gohtml")  // tpl is holding the data i parsed

	// we should use Parse Glob as it stands for a whole directory

	// if err!=nil{
	// 	log.Fatalln(err)
	// }
	err:= tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil) // executing and printing the code
	// if more than one template, execute can run specififc template
	if err!=nil{
		log.Fatalln(err)
	}

}