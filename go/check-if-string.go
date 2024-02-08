var myVar interface{} = "Hello, World!"

if str, ok := myVar.(string); ok {
    fmt.Println("myVar is a string:", str)
} else {
    fmt.Println("myVar is not a string")
}
