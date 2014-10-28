package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("Rscript", "demo_glm/grouping1.R").Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}

/* Show in the slide
// Start OMIT

f1 <- maldib.glm(rings ~ . - id | sex, data = x, family = poisson(log))
f1

// End OMIT
*/
