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
	//verify git, from coip/gitgopher
	if srcgetter, err := exec.LookPath(getIdemSHA.Args[0]); err != nil {
		os.Stderr.Write([]byte(err.Error()))
	} else {
		if *verbose {
			os.Stderr.Write([]byte("plumbing for version: using " + getIdemSHA.Args[0] + "@" + srcgetter))
		}
		getIdemSHA.Path = srcgetter
	}
}

//GetCurrentMainGopher will return a string which attempts to indicate:
// who the gopher is,
// which iteration of the gopher is running, and
// where the gopher is running.
func GetCurrentMainGopher() string {

	//Scope at play:
	var (
		who, where string
		which      []byte //os.Exec output

		err error
	)

	//attempt to identify Who
	who = os.Args[0]

	//attempt to identify Where
	where, err = os.Hostname()
	if err != nil {
		where = "unknownhost"
	}

	//attempt to identify which
	which, err = getIdemSHA.CombinedOutput()

	//Happy path exit!
	if err == nil {
		return where + "/" + who + ":" + string(which)
	}

	//No Which... hmm.
	if *verbose {
		os.Stderr.Write([]byte("[ERROR] pkg idem: failure on exec of [" + getIdemSHA.String() + "]:\n\t" + err.Error() + "\n"))
	}
	which = []byte("unknownSHA")

	//Eh some context is better than none..
	return where + "/" + who + ":" + string(which)

}
