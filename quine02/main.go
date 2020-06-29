// Quine without printf
package main

import "fmt"

func replaceWithItself(s string) (t string) {
	for _, r := range []byte(s) {
		if r == byte(33) {
			t += string(96) + s + string(96)
		} else {
			t += string(r)
		}
	}
	return
}

const s = `// Quine without printf
package main

import "fmt"

func replaceWithItself(s string) (t string) {
	for _, r := range []byte(s) {
		if r == byte(33) {
			t += string(96) + s + string(96)
		} else {
			t += string(r)
		}
	}
	return
}

const s = !

func main() {
	fmt.Print(replaceWithItself(s))
}
`

func main() {
	fmt.Print(replaceWithItself(s))
}
