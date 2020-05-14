# nativify

Create 'native' apps from websites

## Description

nativify is just a wrapper around Chrome application mode and [Desktop entry files](https://specifications.freedesktop.org/desktop-entry-spec/desktop-entry-spec-latest.html).

## Install

To install, use `go get`:

```bash
$ go get -v github.com/someone-stole-my-name/nativify
```

Or download the latest release binary.

## Example app

```bash
nativify add --name Webex --url https://teams.webex.com/ --icon https://www.webex.com/content/dam/wbx/us/images/cisco-webex-teams-icon-180.png
```
