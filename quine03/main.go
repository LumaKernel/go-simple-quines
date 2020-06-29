// Quine not using printf and listing all characters here
// It always can be done beacuse the source code is finite.
// That implies some string exist not appearing in this code.
// abcdefghijklmnopqrstuvwxyz
// ABCDEFGHIJKLMNOPQRSTUVWXYZ
// "/<:=+(){}[] 	.`;
package main

import "fmt"

func replaceWithItself(s string) (t string) {
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i:i+2] == "a"+"a" {
			t += "`" + s + "`"
			i++
		} else if i+1 < len(s) && s[i:i+2] == "b"+"b" {
			t += "`"
			i++
		} else {
			t += s[i : i+1]
		}
	}
	return
}

const s = `// Quine not using printf and listing all characters here
// It always can be done beacuse the source code is finite.
// That implies some string exist not appearing in this code.
// abcdefghijklmnopqrstuvwxyz
// ABCDEFGHIJKLMNOPQRSTUVWXYZ
// "/<:=+(){}[] 	.bb;
package main

import "fmt"

func replaceWithItself(s string) (t string) {
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i:i+2] == "a"+"a" {
			t += "bb" + s + "bb"
			i++
		} else if i+1 < len(s) && s[i:i+2] == "b"+"b" {
			t += "bb"
			i++
		} else {
			t += s[i : i+1]
		}
	}
	return
}

const s = aa

func main() {
	fmt.Print(replaceWithItself(s))
}
`

func main() {
	fmt.Print(replaceWithItself(s))
}
