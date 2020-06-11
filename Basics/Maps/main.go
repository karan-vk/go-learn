package main
// import "fmt"

func main()  {
	// var colors map[string]string
	color:=map[string]string{
	"white":"ffff",
	"Green":"ffff0",
	"yellow":"dmat",
}
println(color)
delete(color,"yellow")
// delete(color,"white")
	
	mapPrint(color)
	
	
}

func mapPrint(c map[string]string)  {
	for color, hex := range c {
		println(color,hex)
	}
	
}