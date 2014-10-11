package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("Rscript", "pivotalr_demo/demo_db_data_frame.R").Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}

/* Show in the slide
// Start OMIT

dat <- db.data.frame('madlibtestdata.dt_abalone') # conn.id = 1, verbose = TRUE

// End OMIT
*/
