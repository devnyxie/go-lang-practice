package main

type Person struct {
	Name string
}

func updateNameByPersonStruct(person *Person) {
	person.Name = "Sasha"

}

func updateNameByType(name *string) {
	*name = "Jane"
}

func updateInnerCopy(name string) {
	nameCopy := name
	println(nameCopy)
	nameCopy = "Jane"
	println(nameCopy)
}

func main() {
	//John
	p := Person{"John"}
	println(p.Name)
	updateNameByPersonStruct(&p)
	println(p.Name)
	//Stacy
	var j string = "Stacy"
	println(j)
	updateNameByType(&j) //pointer to j
	println(j)
	//Bennet
	b := "Bennet"
	println(b)
	updateNameByType(&b) //pointer to j
	println(b)
	//Test
	k := "Test"
	updateInnerCopy(k)
}
