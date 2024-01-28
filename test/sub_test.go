package test_sub

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 这里可以添加一些设置操作

	// 调用所有其他测试函数
	code := m.Run()

	// 这里可以添加一些拆解操作

	// 退出并返回状态码
	os.Exit(code)
}

var name = flag.String("name", "everyone", "The greeting object.")

type Model interface {
	Create() error
	Update() error
	Delete() error
	Select() error
}

type User struct {
	Name string
	Age  int
}

func (u *User) Create() error {
	fmt.Println("create user")
	return nil
}

func (u *User) Update() error {
	fmt.Println("update user")
	return nil
}

func (u *User) Delete() error {
	fmt.Println("delete user")
	return nil
}

func (u *User) Select() error {
	fmt.Println("select user")
	return nil
}

type Dog struct {
	Name string
	Age  int
}

func (d *Dog) Create() error {
	fmt.Println("create dog")
	return nil
}

func (d *Dog) Update() error {
	fmt.Println("update dog")
	return nil
}

func (d *Dog) Delete() error {
	fmt.Println("delete dog")
	return nil
}

func (d *Dog) Select() error {

	fmt.Println("select dog")
	return nil
}

var ModelMap = map[string]Model{
	"user": &User{},
	"dog":  &Dog{},
}

func TestSubtract(t *testing.T) {
	fmt.Printf("Hello, %s!\n", *name)
}

func TestCreate(t *testing.T) {
	ModelMap["user"].Create()
	ModelMap["dog"].Create()
}
