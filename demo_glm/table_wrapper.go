package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("Rscript", "demo_glm/table_wrapper.R").Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}

/* Show in the slide
// Start OMIT

x <- db.data.frame('madlibtestdata.dt_abalone')

// End OMIT
*/
