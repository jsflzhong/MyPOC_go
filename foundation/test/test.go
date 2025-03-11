package main

import (
    "fmt"
    "sync"
    "errors"
    "os"
    "net/http"
)

// 1. 变量与常量
var globalVar int = 10
const PI = 3.14

// 2. 数据类型
type Person struct {
    Name string
    Age  int
}

// 3. 条件语句
func checkAge(age int) {
    if age >= 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }
}

// 4. 循环
func loopExample() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    numbers := []int{1, 2, 3}
    for _, num := range numbers {
        fmt.Println(num)
    }
}

// 5. 函数与方法
func add(a, b int) int {
    return a + b
}

func (p *Person) greet() {
    fmt.Printf("Hello, my name is %s\n", p.Name)
}

// 6. 指针
func pointerExample() {
    x := 10
    p := &x
    *p = 20
    fmt.Println(x)
}

// 7. 接口
type Speaker interface {
    Speak()
}

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Woof!")
}

// 8. 并发编程
func concurrentExample() {
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        defer wg.Done()
        fmt.Println("Goroutine 1")
    }()
    go func() {
        defer wg.Done()
        fmt.Println("Goroutine 2")
    }()
    wg.Wait()
}

// 9. 异常处理
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// 10. 文件操作
func fileExample() {
    file, err := os.Create("test.txt")
    if err != nil {
        fmt.Println("Error creating file")
        return
    }
    defer file.Close()
    file.WriteString("Hello, File!")
}

// 11. 网络编程
func httpExample() {
    resp, err := http.Get("https://www.google.com")
    if err != nil {
        fmt.Println("Request failed")
        return
    }
    defer resp.Body.Close()
    fmt.Println("Status Code:", resp.StatusCode)
}

// 12. 反射
func reflectExample(i interface{}) {
    fmt.Println("Type:", fmt.Sprintf("%T", i))
}

func main() {
    checkAge(20)
    loopExample()
    fmt.Println(add(3, 4))
    person := Person{Name: "Alice", Age: 25}
    person.greet()
    pointerExample()
    dog := Dog{}
    dog.Speak()
    concurrentExample()
    if result, err := divide(10, 2); err == nil {
        fmt.Println("Result:", result)
    }
    fileExample()
    httpExample()
    reflectExample(42)
}
