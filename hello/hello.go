package hello

import (
    // "html/template"
	"fmt"
//	"io"
	"math"
	"math/big"
    "net/http"
)


func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/newtonsmethod", newtonsmethod)
	http.HandleFunc("/sieve", Sieve)	
}


func init_slice(size int) ([]int){
	int_slice:= make([]int, size)
	
	//Range form of for loop does not make sense to me, is this how I use it?
	for i:= range int_slice {
		int_slice[i] = i+2
	}
	
	return int_slice
}

func is_prime(n int) bool {
	n64 := int64(n)
	new_Int := big.NewInt(n64)
	return new_Int.ProbablyPrime(n)
}

func Sieve(w http.ResponseWriter, r *http.Request){
	
	var size int
	size = 100
	
	list_numbers := init_slice(size)
	
	//iterate through 1/2 our list of numbers
	for i:=0; i<(len(list_numbers)/2); i++ {
		//non-zero entries should all be prime as they have no factors
		if list_numbers[i] != 0{
			//remove all subsequent factors for this number
			for j:=i+list_numbers[i]; j<len(list_numbers); j+=list_numbers[i]{
				list_numbers[j] = 0
			}
		}
	}
	
	//Slice should only have primes left in it
	fmt.Fprintln(w, list_numbers)
	
}

func nMethod(x float64, delta float64) (float64, int){

	z := float64(1)
	count := 0
	
	for ; math.Abs((z*z)-x) > delta; {		
			z = z - ( ( (z*z) - x )/(2 * z) )	
			count++
	}	
	return z, count	
}


func newtonsmethod(w http.ResponseWriter, r *http.Request){
	root, count := nMethod(2, .0000000000001)
	fmt.Fprintln(w, "Square root of 2 using Newtons method is: ", root, " cycles: ", count)
	fmt.Fprintln(w, "Square root of 2 is: ", math.Sqrt(2))
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}


