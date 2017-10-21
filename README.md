# gomine

gomine is command line tool for redmine viewer.

# Install

```
go get github.com/hlts2/gomine
```

## Example

Create a ` ~/.config/gomine/config.yaml`

```
url: <please input your redmine url>
apikey: <please input your redmine apikey>
```

## CLI Usage

$ gomine --help
```
Usage:
  gomine [command]

Available Commands:
  cat         Show ticket details of Redmine
  help        Help about any command
  ls          List the Redmine tickets

Flags:
  -h, --help   help for gomine
```

$ gomine cat
```
gomine cat <option> <arguments>

Issues Command:
  cat   show given issue
        $ gomine cat -n 10078 i

Projects Command:
  cat   show given project
        $ gomine cat -n 1 p

Memberships Command:
  cat   show memberships of given project
        $ gomine cat -n 1 m
```

$ gomine ls
```
gomine cat <option> <arguments>

gomine ls <option> <arguments>

Issues Command:
  ls    listing projects
        $ gomine ls i

Projects Command:
  ls    listing projects
        $ gomine ls p
```
