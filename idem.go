package idem

import (
	"flag"
	"os"
	"os/exec"
)

var (
	verbose = flag.Bool("v", false, "exposes some ~wiring for observability")

	getIdemSHA = &exec.Cmd{
		Args: []string{"git", "log", "-1", "--pretty=format:%h"},
	}
)

func init() {
	flag.Parse()
}

//GetCurrentMainGopher will return a string which attempts to indicate:
// who the gopher is,
// which iteration of the gopher is running, and
// where the gopher is running.
func GetCurrentMainGopher() string {

	//attempt to identify Where
	host, err := os.Hostname()
	if err != nil {
		host = "unknown"
	}

	//attempt to identify which
	v, err := getIdemSHA.CombinedOutput()
	if err == nil {
		//Happy path!
		return host + "/" + os.Args[0] + ":" + string(v)
	}

	//No Which... hmm.
	if err != nil && *verbose {
		os.Stderr.Write([]byte("[ERROR] pkg idem: failure on exec of [" + getIdemSHA.String() + "]:\n\t" + err.Error() + "\n"))
	}

	//Eh some context is better than none..
	return host + "/" + os.Args[0] + ":unknownversion"

}
