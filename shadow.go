package main

import "fmt"

type Prov struct {
	Name  string
	Core  string
	Core2 string
}

func (p Prov) Company() Prov {
	fmt.Printf("Компания %s занимается этим %s и %s.\n", p.Name, p.Core, p.Core2)
	return p
}

func (p Prov) Concurent(Name string) {
	fmt.Println("Каждый год соревнуется с", Name)
}

func main() {
	mts := Prov{
		Name:  "MTS",
		Core:  "telecomynication",
		Core2: "IT",
	}

	mts.Company().Concurent("Rosteleckomom")
	Prov.Concurent(Prov.Company(mts), "Rosteleckomom")
}
