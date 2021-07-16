# GitHub top 10 app
It is a test assignment.  
You can find all the preconditions [here](https://github.com/zaffka/analytics-software-engineer-assignment)

The app prints:
* Top 10 active users sorted by pull requests created and commits number
* Top 10 repositories sorted by commits number
* Top 10 repositories sorted by watch events number

It uses some Github stat records collected and compressed in a `data.tar.gz` file.  
(!) The app is using the `embed` package to include that `data.tar.gz` database file right to the binaries at the build stage.
## App structure
```
.
├── data.tar.gz
├── .bin
│   ├── top10-linux
│   ├── top10-mac
├── db
│   ├── db.go
│   ├── db_test.go
│   ├── entities.go
│   ├── helpers.go
│   ├── iterator.go
│   └── iterator_test.go
├── db_scheme.svg
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── README.md
└── stat
    ├── helpers.go
    ├── record.go
    └── stat.go
```
You can find compiled app's binaries at the `.bin` folder.
## Scheme of the datafile
The `data.tar.gz` consists of four main files:  
* events.csv
* actors.csv
* repos.csv
* commits.csv

with such a structure:

![Database scheme](db_scheme.svg)

## How to use the app
Just run the appropriate binary for your OS.  

You can always recompile the app if needed.  
Use `make` command to build a fresh version.