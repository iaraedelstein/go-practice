package main

import "fmt"

func main() {
	fmt.Printf("hola")

	res := Saludar("Iara")

	var name string
	name = "Iara"
	resboolean, resname := Saludar2(name)

	fmt.Println(res)
	fmt.Println(resboolean)
	fmt.Println(resname)

	//slice =array
	//var lista = make([]int, 2)

	Prueba(Saludar2)

	marta:= &Humano{nombre: "MARTA" }
	marta.Correr()
}

func Saludar(name string) bool {
	fmt.Printf("Hola %s", name)
	return true
}

/**
Retorno dos cosas
*/
func Saludar2(name string) (bool, string) {
	fmt.Printf("Hola %s", name)
	return true, name
}

/**
Recibo y retorno puntero humano
*/
func Saludar3(h *Humano) (bool, *Humano) {
	fmt.Printf("Hola %s", h.nombre)
	return true, h
}

func Prueba(p Persona)  (bool, string){
	return p("Pepito")
}

type Persona func( nome string) (bool, string)


type Humano struct {
	nombre string
	Runner
}

type Runner struct {

}

//Asigne una funcion a un struct
func (c * Runner) Correr()  bool {
	return true
}

//interface
//verbo en infinitivos serilizer
type Corredor interface {
	Correr()
}
