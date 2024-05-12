# Documentation

## FIRST RULE

Each folder containing in cmd folder is a top of level command (excludig the root command).
if the file.go has the same name of the folder then the file.go is the root command of the palet.

## SECOND RULE

Each file.go is a main palet that contains the main commands that display their options and purpose,
and a function that add their subcommands.

Following that logic we can implement a command line interface that is easy to understand and to maintain
with several levels of nested commands.
