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

gp <- db.connect(port=16526, dbname='madlib', host='127.0.0.1', user='gpadmin') # 1, Greenplum database
pg <- db.connect(port=5333, dbname='madlib', host='127.0.0.1', user='gpadmin')  # 2, Postgres database
hq <- db.connect(port=18526, dbname='madlib', host='127.0.0.1', user='gpadmin') # 3, HAWQ database
db.list()

// End OMIT
*/
