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

```
gomine <command> <option> <arguments>

Issues Command:
  ls    listing issue
        $ gomine ls i

  cat   show given issue
        $ gomine cat -n 10078 i

Projects Command:
  ls    listing projects
        $ gomine ls p

  cat   show given project
        $ gomine cat -n 1 p

Memberships Command:
  cat   show memberships of given project
        $ gomine cat -n 1 m

```
