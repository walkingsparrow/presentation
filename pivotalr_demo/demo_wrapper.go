package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("Rscript", "pivotalr_demo/demo_wrapper.R").Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}

/* Show in the slide
// Start OMIT

dat

// End OMIT
*/
