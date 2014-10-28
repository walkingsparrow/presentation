package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("Rscript", "demo_glm/connect.R").Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}

/* Show in the slide
// Start OMIT

gp <- db.connect(port=16526, dbname='madlib', host='127.0.0.1', user='gpadmin') # 1, GPDB
db.list()

// End OMIT
*/
