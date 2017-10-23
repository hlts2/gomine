# gomine

gomine is command line tool for viewing Redmine.

# Install

```
go get github.com/hlts2/gomine
```

## Example

Create a `config.yaml`

#### Windows
```
C:\> type nul > C:\Users\<username>\AppData\.config\gomine\config.yaml
```

#### Unix
```
$ touch ~/.config/gomine/config.yaml
```

Input your Redmine info

```
url: <Redmine URL>
apikey: <Redmine APIKEY>
```

## CLI Usage


```
$ gomine --help
Usage:
  gomine [command]

Available Commands:
  cat         Show "issues", "projects", "memberships" details of Redmine
  help        Help about any command
  ls          List "issues", "projects", "memberships" of Redmine

Flags:
  -h, --help   help for gomine
```

```
$ gomine cat
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

```
$ gomine ls
gomine ls <option> <arguments>

Issues Command:
  ls    listing projects
        $ gomine ls i

Projects Command:
  ls    listing projects
        $ gomine ls p
```
