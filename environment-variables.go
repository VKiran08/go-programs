package main

import "os"  // This is teh OS package. Go thru the documentation for so many functions
import "fmt"  //



func main(){

fmt.Println("All environment variables  ---->", os.Environ())
fmt.Println("Username is ---->", os.Getenv("USER"))
fmt.Println("User's home directory name is ---->", os.Getenv("HOME"))

}
 
