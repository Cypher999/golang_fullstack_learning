package main
import "fmt"

type Worker interface{
	Work() string
}

type Programer struct{
	name string
	prefered_language string
	project_count int
}

type Designer struct{
	name string
	prefered_tools string
	project_count int
}

func (p Programer) Work() string{
	return fmt.Sprintf("%s is a programer, he/she prefer %s for programing language, so far he/she has complete %d project/s",p.name,p.prefered_language,p.project_count)
}

func (d Designer) Work() string{
	return fmt.Sprintf("%s is a designer, he/she prefer %s for design tools, so far he/she has complete %d project/s",d.name,d.prefered_tools,d.project_count)
}

func main(){
	worker_sandi:=Programer{name:"Sandi Mhd Irvan",prefered_language:"php, javascript, python",project_count:10}
	worker_deja:=Designer{name:"Deja Almustaqim",prefered_tools:"adobe photoshop, canva",project_count:10}
	fmt.Println(worker_sandi.Work())
	fmt.Println(worker_deja.Work())
}