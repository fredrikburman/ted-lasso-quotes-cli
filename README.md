Ted Lasso Quotes CLI
====================
This cli makes it easy to get quotes from the [Ted Lasso Quotes API](https://tedlassoquotes.com) 

Just clone the repo and build the binary for your distribution

Prerequisite: you need to have Go 1.15 or later installed

This should be enough to build your binary `go build -o tedlasso-cli main.go`

Then run it using
`./tedlasso-cli`

## Available params
``` 
./tedlasso-cli -h
  -c string
    	Character to fetch (default to random: valid choices: ted-lasso|roy-kent|coach-beard)
  -o string
    	Print output in format: text/json (default "text")
  -t duration
    	Client timeout (default 2s)
```

## Example

```
$ ./tedlasso-cli -c roy-kent

Enjoy your trophies for winning nothing.”
— Roy Kent
```
