package main


import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type calculus struct {
	First float64 
	Operation string 
	Second float64
	Igual string
	Result float64 

} 

 var history = []calculus{}

func hanleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/calc", calc)
	myRouter.HandleFunc("/calc/history", getHistory)
	myRouter.HandleFunc("/calc/sum/{a}/{b}", sum)
	myRouter.HandleFunc("/calc/sub/{a}/{b}", sub)
	myRouter.HandleFunc("/calc/mult/{a}/{b}", mult)
	myRouter.HandleFunc("/calc/div/{a}/{b}", div)

	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func  calc( w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Calculator!")
}

func sum( w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	a := vars["a"]
	b := vars["b"]
	num1, err1 := strconv.ParseFloat(a, 64) 
	num2, err2 := strconv.ParseFloat(b, 64)
	if (err1 == nil && err2 == nil) {
		result := num1 + num2
		history =  append(history, calculus{num1,"+", num2,"=" , result})
		fmt.Fprintf(w, "Sum = %.2f", result)
	} else {
		fmt.Fprint(w,"invalid number")
	}
}

func sub( w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	a := vars["a"]
	b := vars["b"]
	num1, err1 := strconv.ParseFloat(a, 64) 
	num2, err2 := strconv.ParseFloat(b, 64)
	if (err1 == nil && err2 == nil) {
		result := num1 - num2
		history =  append(history,calculus{num1,"-", num2,"=" , result})
		fmt.Fprintf(w, "Sub = %.2f", result)
	} else {
		fmt.Fprint(w,"invalid number")
	}
}

func mult( w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	a := vars["a"]
	b := vars["b"]
	num1, err1 := strconv.ParseFloat(a, 64) 
	num2, err2 := strconv.ParseFloat(b, 64)
	if (err1 == nil && err2 == nil) {
		result := num1 * num2
		history =  append(history,calculus{num1,"*", num2,"=" , result})
		fmt.Fprintf(w, "Mult = %.2f", result)
	} else {
		fmt.Fprint(w,"invalid number")
	}
}

func div( w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	a := vars["a"]
	b := vars["b"]
	num1, err1 := strconv.ParseFloat(a, 64) 
	num2, err2 := strconv.ParseFloat(b, 64)
	if (err1 != nil || err2 != nil) {  
		fmt.Fprint(w, "Invalid number. Please try again! ")
	} else if ( num2 == 0 ){
		fmt.Fprint(w, "Cannot divide by zero. Please try again! ")
	} else {
		result := num1 / num2
		history =  append(history,calculus{num1,"/", num2,"=" , result})
		fmt.Fprintf(w, "Div = %.2f", result)
	}
}

func getHistory( w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, history)
}

func main() {
	fmt.Print("Calculator Running http://localhost:8082/calc")
	hanleRequests()	
}
