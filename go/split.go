import (
    "fmt"
    "strings"
)

func main() {
    arn := "a:p:r:a:s:a:b:b:c:r"
    
    // Splitting the string into a slice
    parts := strings.Split(arn, ":")

    // Creating new strings by concatenating the required parts
    if len(parts) >= 6 {
        a := parts[4] + ":" + parts[5]
        b := parts[6]

        fmt.Println("ID1:", a)
        fmt.Println("ID2:", b)
    } else {
        fmt.Println("string does not contain enough parts")
    }
}
