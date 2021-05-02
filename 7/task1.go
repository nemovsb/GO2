package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"
)

type in struct {
	Name    string
	Surname string
	Age     int
	Height  int
	Weight  int
}

type values map[string]interface{}

func MapToStruct(myStruct *in, myMap values) error {

	if myStruct == nil {
		return errors.New("arg is nil")
	}

	valStruct := reflect.ValueOf(myStruct)
	if valStruct.Kind() == reflect.Ptr {
		valStruct = valStruct.Elem()
	}

	for i := 0; i < valStruct.NumField(); i++ {
		typeFieldName := valStruct.Type().Field(i).Name

		//log.Printf("reflect.ValueOf(myMap[typeFieldName]): %v\n", reflect.ValueOf(myMap[typeFieldName]))
		//log.Printf("typeFieldName: %s", typeFieldName)

		valStruct.Field(i).Set(reflect.ValueOf(myMap[typeFieldName]))
	}

	return nil
}

func main() {
	var myStruct in

	var myMap = values{
		"Name":    "Sergey",
		"Surname": "Nemov",
		"Age":     28,
		"Height":  190,
		"Weight":  90,
	}
	//log.Printf("myStruct: %+v\n", myStruct)
	MapToStruct(&myStruct, myMap)
	log.Printf("myStruct: %+v\n", myStruct)
	//fmt.Println(myStruct)
}

func ExampleMapToStruct() {
	var myStruct in
	var myMap = values{
		"Name":    "Sergey",
		"Surname": "Nemov",
		"Age":     28,
		"Height":  190,
		"Weight":  90,
	}
	MapToStruct(&myStruct, myMap)
	fmt.Printf("myStruct: %+v\n", myStruct)
	//Output: myStruct: {Name:Sergey Surname:Nemov Age:28 Height:190 Weight:90}
}

func TestCalcNumberWithoutRecursion(t *testing.T) {
	testSets := []struct {
		name     string
		myMap    values
		myStruct in
		wait     in
	}{
		{
			name: "1",
			myMap: values{
				"Name":    "Sergey",
				"Surname": "Nemov",
				"Age":     28,
				"Height":  190,
				"Weight":  90,
			},
			wait: in{
				Name:    "Sergey",
				Surname: "Nemov",
				Age:     28,
				Height:  190,
				Weight:  90,
			},
		},
	}
	for _, get := range testSets {
		t.Run(get.name, func(t *testing.T) {
			t.Log(get.myMap)
			if MapToStruct(&get.myStruct, get.myMap); get.myStruct != get.wait {
				t.Errorf("Got %+v, but wait %+v\n", get.myStruct, get.wait)
			}
		})
	}
}
