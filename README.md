# Exercise resolution

## How to compile
On the root of the project (where `main.go` is), run:
> go build

## How to run 
An executable `secscan` should be created in the same directory after successful compilation.
> ./secscan scan ./testproject

The command above will run a scan on folder `testproject/` using the default config found 
on the root of the project

For help on how to run:
> ./secscan -h
> ./secscan scan -h

### Custom config
It is possible to provide a custom configuration to the application.
To do so please refer to the default config file `config.yml`.
<br>Make a copy of it and change it at will (well, careful not to break :smile: )

To execute the application using the custom config:
> ./secscan scan -c ./config.yml ./testproject

It's also possible to override the output format specified in the config file, use `-j`:
 > ./secscan scan -j -c ./config.yml ./testproject

#### Possible configs
It is possible to provide several terms for the sensitive data exposure verification.
Just add them to:
```yml
  sensitiveTerms:
    - ACME
    - Tom & Jerry
    - $1.15b
```
Specify which extensions should be targeted by the x-site scripting checkers:
```yml
  xSSTargetExtensions:
    - .js
    - .html
```

Exclude checkers from being used:
```yml
  excludedCheckers: 
    - SDEChecker
    - SQLInjectionChecker
    - XSSChecker
``` 

Exclude files from being checked
```yml
  excludedFileExtensions:
    - .yml
```

## Extending the report modes
To add new report modes just implement `Writer interface` and register it in the `delegate.go` file.

## Extending the checkers available
To add new checkers just implement `hecker interface`  then register it in the `scan/scan.go` file.
```go
var checkers = []Checker{
	SQLInjectionChecker{},
	XSSChecker{},
	SDEChecker{},
}
```

# Final notes
## Tests
Unfortunately could not test everything

## sqlinjection.go
Did not have time to finish all the scenarios I had planned:
* multi line sql queries 

## Packages
All packages inside `pkg` should be reusable, that is why they do not depend on each other
* separate configs
* to each their own model

# Exercise Description
Create a console application that mimics a security code scanner.
<br>The application accepts path to the source code and scan configuration in the parameters, performs
security analysis and prints report to output.

## Checks
There are 3 types of security checks and adding the additional types should be simple.

### Cross site scripting
Check that there are no html or JavaScript files with following statement:
> Alert()

### Sensitive Data Exposure
Check that there are no files with following strings on a same line:
> “ACME” “Tom & Jerry” “$1.15b”

### SQL Injection
Check that there are no files
with statements starting with quotes, containing SELECT, WHERE, %s and ending with quotes, e.g:
> ".... SELECT .... WHERE .... %s .... "

## Results
There are 2 supported output formats and adding the additional formats should be simple.
- Plain text with vulnerability per line, e.g:
  - > [SQL injection] in file “DB.go” on line 45
- Json representing the same information

