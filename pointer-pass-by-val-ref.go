package main

import "fmt"

func main(){
a:= 10
b:= 20
//declare p as a pointer of b - & gives the address and * gives the value of b
p:= &b
fmt.Println("P address befoew is ", p)
fmt.Println("P values before is ", *p)
b=21
fmt.Println("P address after is ", p)
fmt.Println("P values after is ", *p)

//pass by value - even though "a" passed to function hello1, "a" value dosnt change in this function
c, d :=passvalue(a)
fmt.Println("A value after",  a)
fmt.Println("Returned values from function passvalue are ",  c, d)

//pass by reference - variable b being passed as reference , so any changes done the called function will effect the value of b here
fmt.Println("B value Before",  b)
e, f :=passreference(p)   //pass the pointer varibale for "b"
fmt.Println("B value after",  b)
fmt.Println("Returned values from function passreference are ",  e, f)
}

//function for pass by value
func  passvalue(a int  ) (x int, y int) {
a=100    //This change deosnt effect the variable a in the calling fucntion - Thi sis called pass by value
x=a+10
y=a+11
return x, y
} 

//function for pass by reference 
func  passreference(a *int  ) (x int, y int) {   //notice the asterisk before the type 
*a=100    //This will effect the variable "b" in the calling function - This is called pass by reference
x=*a+10
y=*a+11
return x, y
} 
