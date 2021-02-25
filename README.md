# TOKOIN CODE CHALLENGE

## Requirements

Using the provided data (**tickets.json** and **users.json** and **organization.json**)

Write a simple command line application to search the data and return the results
in a human readable format.

* Feel free to use libraries or roll your own code as you see fit.
* Where the data exists, values from any related entities **MUST** be included in
the results, i.e.
    * Searching **organization MUST** return its **ticket subject** and **users name**.
    * Searching **users MUST** return his/her **assignee ticket subject** and **submitted ticket subject** and his/her **organization name**.
    * Searching **tickets MUST** return its **assignee name**, **submitter name**, and **organization name**.
* The user should be able to search on any field, full value matching is fine
(e.g. "mar" won't return "mary").
* The user should also be able to search for empty values, e.g. where
description is empty.
* Search can get pretty complicated pretty easily, we just want to see that you
can code a basic search application.

## Introduction

A simple test challenge from tokoin.

#### Setup
* Install require packages: `go mod vendor`

#### Build and Run
##### Run
```shell script
go run main.go
```
##### Test
```shell script
go test ./src/tests -v
```

#### Structure
```shell script
├── config              # Contain config file
├── src                 # Source code
│   └── common          # Common function 
│   └── data            # Contain data files
│   └── entity          # Entity Layer: contain entity
│   └── repositories    # Repositories Layer: contain interface and implement
│   └── services        # Business Layer: contain interface and logic implement
├── tests               # Tests package: defaine testcases and implement unittest
│   └── data            # Contain data test files
```

### Happy coding :))