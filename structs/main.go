package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact contactInfo
	contactInfo //easier way to declare a field
}

func main() {

	//primer forma de construir una persona
	//iara := person{ "iara", "edelstein"}

	// esta forma es más segura
	//iara := person{ firstName: "iara", lastName: "edelstein"}
	//fmt.Println(iara)

	//tercera forma
	// var iara person
	// iara.firstName = "iara"
	// iara.lastName = "edelstein"
	// fmt.Println(iara)
	// fmt.Printf("%+v", iara)

	jim := person{
		firstName: "jim",
		lastName:  "yang",
		contactInfo: contactInfo{
			email:   "holajim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.print()

	//pointers
	//jimPointer := &jim
	//jimPointer.updateName("jimmy")

	//simplificado los receivers automaticamente convierten a puntero de ser necesario
	jim.updateName("jimmy")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}


//&variable : operator that gives the access to the memory address where the variable is located
//*pointer : give me de value that exist on that memory address
func (p *person) updateName( name string) {
	(*p).firstName = name
}
