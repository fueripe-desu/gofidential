package main

import (
	"fmt"
	"log"
	"time"

	gf "github.com/fueripe-desu/gofidential/v2"
)

type Secrets struct {
	IntField       int
	SmallIntField  int8
	ShortIntField  int16
	NormalIntField int32
	LargeIntField  int64

	UintField       uint
	SmallUintField  uint8
	ShortUintField  uint16
	NormalUintField uint32
	LargeUintField  uint64

	FloatField      float32
	LargeFloatField float64

	ComplexField      complex64
	LargeComplexField complex128

	StringField string
	BoolField   bool
	TimeField   time.Time
}

func main() {
	env := gf.Environment{Name: "dev"}
	s := Secrets{}

	err := gf.Load(&s, env)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Int field: '%v'\n", s.IntField)
	fmt.Printf("Small int field: '%v'\n", s.SmallIntField)
	fmt.Printf("Short int field: '%v'\n", s.ShortIntField)
	fmt.Printf("Normal int field: '%v'\n", s.NormalIntField)
	fmt.Printf("Large int field: '%v'\n", s.LargeIntField)

	fmt.Printf("Uint field: '%v'\n", s.UintField)
	fmt.Printf("Small uint field: '%v'\n", s.SmallUintField)
	fmt.Printf("Short uint field: '%v'\n", s.ShortUintField)
	fmt.Printf("Normal uint field: '%v'\n", s.NormalUintField)
	fmt.Printf("Large uint field: '%v'\n", s.LargeUintField)

	fmt.Printf("Float field: '%v'\n", s.FloatField)
	fmt.Printf("Large float field: '%v'\n", s.LargeFloatField)

	fmt.Printf("Complex field: '%v'\n", s.ComplexField)
	fmt.Printf("Large complex field: '%v'\n", s.LargeComplexField)

	fmt.Printf("String field: '%v'\n", s.StringField)
	fmt.Printf("Bool field: '%v'\n", s.BoolField)
	fmt.Printf("Time field: '%v'\n", s.TimeField.String())
}
