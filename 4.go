package main
import (
  "fmt"
)
func  main(){
  var n int

  fmt.Println("Enter length of numbers:")
  fmt.Scan(&n)
  a := make([]int, n)
  fmt.Println("Enter the inputs:")
  for i:=0;i<n;i++{
    fmt.Scan(&a[i])
  }
  fmt.Println("Before Swapping:",a)
  for i:=0;i<n/2;i++{
    swap(&a[i],&a[n-(i+1)])
  }
  fmt.Println("After swapping", a)
}
func swap(a, b *int){
  *a,*b = *b,*a
  fmt.Println(*a,*b)
}