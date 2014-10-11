package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("Rexec", "pivotalr_demo/demo_connect.R").Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}

/* Show in the slide
// Start OMIT

pg <- db.connect(port = 5333, dbname = 'madlib')  # 1, Postgres database
gp <- db.connect(port = 16526, dbname = 'madlib') # 2, Greenplum database
hq <- db.connect(port = 18526, dbname = 'madlib') # 3, HAWQ database
db.list()

// End OMIT
*/
