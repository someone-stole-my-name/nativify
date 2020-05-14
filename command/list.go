package command

import (
	"io/ioutil"
	"log"
	"os/user"
	"regexp"
	"strings"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
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
			c.Ui.Output(name[1])
		}
	}

	return 0
}

func (c *ListCommand) Synopsis() string {
	return "List nativified apps"
}

func (c *ListCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
