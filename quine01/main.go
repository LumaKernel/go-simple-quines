// Quine with printf
package main

import "fmt"

const s = `// Quine with printf
package main

import "fmt"

const s = %s

func main() {
	fmt.Printf(s, string(96)+s+string(96))
}
`

func main() {
	fmt.Printf(s, string(96)+s+string(96))
}
