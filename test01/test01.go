package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("Rscript", "test01/test01.R").Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))

	out, err = exec.Command("skim", "test01.pdf").Output()
	if err != nil {
		fmt.Println(err)
		return
	}
}

/* Show the following in the slides
// Start OMIT

library(PivotalR)

f <- lm(rings ~ . - id, data = abalone)

summary(f)

plot(abalone$rings, predict(f))

// End OMIT
*/
