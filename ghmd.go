package ghmd

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/go-github/github"
)

const cmdName = "ghmd"

// Run the ghmd
func Run(argv []string, outStream, errStream io.Writer) error {
	log.SetOutput(errStream)
	fs := flag.NewFlagSet(
		fmt.Sprintf("%s (v%s rev:%s)", cmdName, version, revision), flag.ContinueOnError)
	fs.SetOutput(errStream)
	ver := fs.Bool("version", false, "display version")
	if err := fs.Parse(argv); err != nil {
		return err
	}
	if *ver {
		return printVersion(outStream)
	}

	md, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	cl := github.NewClient(nil)
	html, _, err := cl.Markdown(context.Background(), string(md), nil)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(outStream, html)
	return err
}

func printVersion(out io.Writer) error {
	_, err := fmt.Fprintf(out, "%s v%s (rev:%s)\n", cmdName, version, revision)
	return err
}
