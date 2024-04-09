package main

import "fmt"

func main(){

    var a,b,c= 2,"valdeilton",6.8

    fmt.Println(a,b,c);

    var(
    	idade=40
    	nome="valdeilton"
    	cpf="00020002002"
    	rua="Jose pessoa"
    )

    fmt.Println("\n Pessoa ",nome,"\n Idade",
    	idade,"\n cpf ",cpf,"\n rua ",rua);

    const PI =3.14
    fmt.Println(PI);

    /*Go Output Functions*/
    //Print
    //Println
    //Printf
		var i,j string ="Hello","Word"
		fmt.Println(i)
		fmt.Println(j)

	/*
	The Printf() Function
	*/

	fmt.Printf("Qual o tipo da constate %T",PI)
	fmt.Printf("\nQual o tipo da constate %%T",PI)
	fmt.Printf("\nQual o valor da variavel %#v\n",j)

  /*
    %b	Base 2
	%d	Base 10
	%+d	Base 10 and always show sign
	%o	Base 8
	%O	Base 8, with leading 0o
	%x	Base 16, lowercase
	%X	Base 16, uppercase
	%#x	Base 16, with leading 0x
	%4d	Pad with spaces (width 4, right justified)
	%-4d	Pad with spaces (width 4, left justified)
	%04d	Pad with zeroes (width 4

  */


       
}