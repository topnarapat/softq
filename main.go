package main

import "fmt"

type Square struct {
	width  uint
	heigth uint
}

type User struct {
	Email string
}

func main() {
	x := 10
	y := &x
	*y = 20
	fmt.Printf("Address x = %v\n", &x)
	fmt.Printf("y = %v\n", y)

	u := Square{}
	u.width = 5
	u.heigth = 4
	z := &u
	fmt.Printf("%v\n", u)
	fmt.Printf("%v\n", z)

	result := Area(u)
	fmt.Printf("%v\n", result)

	user := SetUser("Chopper")
	fmt.Printf("%v\n", user)
	fmt.Printf("%v\n", user.Email)
}

func Area(a Square) uint {
	return a.width + a.heigth
}

//* ในกรณีที่ func จะ return struct โดยอาจจะมีการเช็คค่าแล้วต้อง return nil จะต้อง return struct เป็น pointer struct
func SetUser(email string) *User {
	if email == "" {
		//* มีการ return nil จึงต้องเป็น *User
		return nil
	}

	//* เมื่อ return เป็น *User ซึ่งก็คือ pointer struct จึงจำเป็นต้องใช้ &User{}
	u := &User{}
	u.Email = email
	return u
}

func SetUserEmail(email string) User {
	//* เมื่อ return เป็น User จึงไม่จำเป็นต้องใช้ &User{}
	u := User{}
	u.Email = email
	return u
}
