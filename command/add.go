package command

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path"
	"strings"
)

type App struct {
	Name, Exec, Icon, StartupWMClass string
}

type AddCommand struct {
	Meta
}

func (c *AddCommand) Run(args []string) int {
	if len(args) != 3 {
		fmt.Println(c.Help())
		return 1
	}

	iconpath, err := ensureIcon(args[2])
	if err != nil {
		panic(err)
	}

	appTemplate, _ := template.New("App").Parse("[Desktop Entry]\nVersion=1.0\nType=Application\nName={{.Name}}\nComment=Nativify\nExec={{.Exec}}\nIcon={{.Icon}}\nTerminal=false\nStartupNotify=true\nStartupWMClass={{.StartupWMClass}}")
	var (
		buf bytes.Buffer
		app = App{
			Name:           args[0],
			Exec:           "google-chrome --new-window --app=" + args[1],
			Icon:           iconpath,
			StartupWMClass: genWMClass(args[1]),
		}
	)
	if err := appTemplate.Execute(&buf, app); err != nil {
		panic(err)
	}

	usr, _ := user.Current()
	dir := usr.HomeDir
	err = ioutil.WriteFile(dir+"/.local/share/applications/"+strings.ToLower(strings.Replace(args[0], " ", "", -1))+".desktop", buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	c.Ui.Output("Done")
	return 0
}

func (c *AddCommand) Synopsis() string {
	return "Nativify an app"
}

func (c *AddCommand) Help() string {
	helpText := `
Usage: add --url string --icon string --name string

url: Web App to nativify
icon: Url to download icon from
name: App Name
`
	return strings.TrimSpace(helpText)
}

func genWMClass(k string) string {
	u, err := url.Parse(k)
	if err != nil {
		panic(err)
	}
	return u.Hostname()
}

func ensureIcon(u string) (icon string, err error) {
	ur, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	filename := path.Base(ur.Path)
	usr, _ := user.Current()
	dir := usr.HomeDir
	resp, err := http.Get(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	out, err := os.Create(dir + "/.local/share/icons/" + filename)
	if err != nil {
		return
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	icon = dir + "/.local/share/icons/" + filename
	return
}
