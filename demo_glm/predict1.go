package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("Rscript", "demo_glm/predict1.R").Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}

/* Show in the slide
// Start OMIT

p <- predict( f )
lk( cbind( x$rings, p ), 10 )

// End OMIT
*/
