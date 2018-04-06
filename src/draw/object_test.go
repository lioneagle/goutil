package draw

import (
	"fmt"
	"testing"
)

type A struct {
	x int
}

func (this *A) f1() {
	fmt.Println("A.f1()")
}

func (this *A) f2() {
	fmt.Println("A.x =", this.x)
}

type B struct {
	A
}

func (this *B) f1() {
	fmt.Println("B.f1()")
}

type C struct {
	A
}

type D struct {
	B
}

func (this *D) f1() {
	fmt.Println("D.f1()")
}

type I interface {
	f1()
}

func TestBase(t *testing.T) {
	a := &A{}
	b := &B{}
	c := &C{}
	d := &D{}

	var i I
	i = a
	i.f1()
	i = b
	i.f1()
	b.A.x = 1
	d.x = 123
	i = c
	i.f1()
	i = d
	i.f1()
	d.B.f1()
	fmt.Println("d.x =", d.x)
	fmt.Println("d.A.x =", d.A.x)
	fmt.Println("d.B.x =", d.B.x)

	v1, ok := i.(*A)
	fmt.Println("ok =", ok)
	fmt.Println("v1 =", v1)

	d.x = 13

	var a2 *A
	d.A.f2()

	a2 = &d.A
	a2.f2()
	a2.f1()

	x1, ok := i.(*A)
	fmt.Println("&ok =", &ok)
	fmt.Println("ok =", ok)
	fmt.Println("x1 =", x1)

	x2, ok := i.(*D)
	fmt.Println("&ok =", &ok)
	fmt.Println("ok =", ok)
	fmt.Println("x2 =", x2)

	var u1 uintptr
	u1 = 1

	u2 := int64(u1)

	fmt.Println("u1 =", u1)
	fmt.Println("u2 =", u2)

	u2 = 100
	u1 = uintptr(u2)

	fmt.Println("u1 =", u1)
	fmt.Println("u2 =", u2)

	object := &Object{}
	var iobect IObject

	iobect = object
	fmt.Println("iobect.GetName() =", iobect.GetName())

}
