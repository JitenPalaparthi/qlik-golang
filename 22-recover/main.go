package main

import "fmt"

type any = interface{}

func main() {
	// fmt.Println("Divid by zero", 10/0)
	//v := 0
	/*defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()*/
	defer fmt.Println("Hello World ends!")
	fn := new(string)
	*fn = "Jiten"
	func() {
		fullName(nil, nil)
		defer fmt.Println("Divid by zero end")
		//fmt.Println("Divid by zero start")
		//fmt.Println("Divid by zero", 10/v)
		//fmt.Println("Checking; Does this work becasue of panic in fullName?")
	}()

	fmt.Println("Hello World starts!")
}
func recoverFullName() {
	if r := recover(); r != nil {
		switch r.(type) {
		case PanicData:
			pd := r.(PanicData)
			fmt.Println("Data:", pd.Data, "Error", pd.Error)
			//fmt.Println("Recovering from fullname")
			//fmt.Println(r)
		default:
			fmt.Println(r)
		}
	}
}

func fullName(fn, ln *string) {
	defer recoverFullName()
	defer println("end calling fullName")
	if fn == nil {
		panic("firstname is nil")
	}
	if ln == nil {
		panic(PanicData{Data: *fn, Error: "ln is nil"})
	}
	println(*fn + " " + *ln)
}

type PanicData struct {
	Data  interface{}
	Error string
}
