# jiratime [![Build Status](https://travis-ci.org/baniol/jiratime.svg?branch=master)](https://travis-ci.org/baniol/jiratime)

[![Coverage Status](https://coveralls.io/repos/github/baniol/jiratime/badge.svg?branch=master)](https://coveralls.io/github/baniol/jiratime?branch=master)

Small CLI utility for displaying logged jira hours.

## Usage

### `jiratime days`

Displays a list of working days of a given month with corresponding logged hours

Example output:

```
Day             Hours
2017-08-01      0
2017-08-02      2
2017-08-03      0
2017-08-04      0
2017-08-07      3
2017-08-08      0
2017-08-09      0
2017-08-10      4
2017-08-11      7
2017-08-14      0
2017-08-15      0
2017-08-16      7
2017-08-17      4
2017-08-18      0
2017-08-21      7
2017-08-22      7
2017-08-23      7
2017-08-24      7
2017-08-25      7
2017-08-28      0
2017-08-29      0
2017-08-30      0
2017-08-31      0
-----------------
Logged  Required
62      161
```

#### Paramteters

* `-year` - optional (current year if not specified)

* `-month` - optional (current month if not specified)

### `jiratime tickets`

Displays a list of tickets with corresponding logged hours

Example output:

```
Ticket  Hours
TK-182  46
TK-59   14
TK-8    8
TK-223  16
----------------------
Total logged: 84 hours
```

## Configiration

There must be a file `jiratimeconfig.yml` present in the user's home directory (see the `config.yml.template` for reference):

```
jiraurl: ""
jirauser: ""
jirapassword: ""
datefrom : "2017-03-01"
hoursdaily: 7
maxsearchresults: 1000
```
