package main

import "fmt"

type Phone interface {
	call() string
}

type Android struct {
	brand string
}
type Iphone struct {
	version string
}

func (android Android) call() string {
	return "I am Android " + android.brand
}

func (ihpone Iphone) call() string {
	return "I am Iphone " + ihpone.version
}

func printcall(p Phone) {
	fmt.Println(p.call() + ",我们都能打电话")
}
func main() {
	var vivo = Android{brand: "Vivo"}
	var hw = Android{brand: "HuaWei"}
	i7 := Iphone{version: "7 plus"}
	ix := Iphone{version: "X"}

	printcall(vivo)
	printcall(hw)
	printcall(i7)
	printcall(ix)

}
