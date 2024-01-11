package test_sub

import (
	"flag"
	"fmt"
	"testing"
	"os"
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

func TestSubtract(t *testing.T) {
	fmt.Printf("Hello, %s!\n", *name)
}
