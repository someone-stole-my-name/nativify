package command

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"regexp"
	"strings"
)

type DeleteCommand struct {
	Meta
}

func (c *DeleteCommand) Run(args []string) int {
	usr, _ := user.Current()
	dir := usr.HomeDir
	root := dir + "/.local/share/applications/"
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		b, err := ioutil.ReadFile(dir + "/.local/share/applications/" + file.Name())
		if err != nil {
			panic(err)
		}
		if strings.Contains(string(b), "Nativify") {
			pat := regexp.MustCompile(`Name=.*`)
			ext := pat.FindString(string(b))
			name := strings.Split(ext, "=")
			if name[1] == args[0] {
				fmt.Println(file.Name())
				err := os.Remove(dir + "/.local/share/applications/" + file.Name())
				if err != nil {
					panic(err)
				}
				c.Ui.Output("Deleted " + args[0])
				return 0
			}
		}
	}
	c.Ui.Warn("Cannot find: " + args[0])
	return 1
}

func (c *DeleteCommand) Synopsis() string {
	return "Remove a nativified app"
}

func (c *DeleteCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
