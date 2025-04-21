# Pointers
What is a pointer? A pointer is a memory address that points to another memory address which stores a value.

pointers in GO are used as follows as follows

```golang
func main() {
    i, j = 42, 2701
    p := &i         // Points to the memory address of i
    *p = 21         // Set the value of i through the pointer
    p = &j          // Points to the memory address of i
    *p = *p / 37    // Divides the value of j by 37
}
```