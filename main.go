package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "deployment, d",
			Value: "deployment",
			Usage: "name of the deployment you want to patch",
		},
	}

	// patch deployment console-deployment -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"creationTimestamp\":\"`date +'%s'`\"}}}}}"

	app.Action = func(c *cli.Context) error {
		deployment := ""
		if c.NArg() == 0 {
			fmt.Printf("No argument detected. use like 'kpatch api-deployment'\n")
			return nil
		}
		if c.NArg() > 0 {
			deployment = c.Args().Get(0)
		}
		kubeCommand := "kubectl patch deployment " + deployment + " -p \"{\\\"spec\\\":{\\\"template\\\":{\\\"metadata\\\":{\\\"labels\\\":{\\\"creationTimestamp\\\":\\\"`date +'%s'`\\\"}}}}}\""
		fmt.Printf(" executing %s ", kubeCommand)
		out, err := exec.Command("bash", "-c", kubeCommand).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The date is:\n%s", out)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
