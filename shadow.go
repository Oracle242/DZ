package main

import "fmt"

type SmarPtpone struct {
	Brand string
	Model string
	OS    string
}

func (s SmarPtpone) Praise() {
	fmt.Printf("Самый надежный телефон %s %s на базе операционной системмы %s.\n", s.Brand, s.Model, s.OS)
}

func main() {
	apple := SmarPtpone{
		Brand: "Apple",
		Model: "13 pro Max",
		OS:    "iOS",
	}

	samsung := SmarPtpone{
		Brand: "Samsung",
		Model: "Pipulia",
		OS:    "Android",
	}
	apple.Praise()
	SmarPtpone.Praise(samsung)
}
