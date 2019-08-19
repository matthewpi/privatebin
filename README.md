# PrivateBin CLI
A CLI for PrivateBin allowing easy pasting from the Terminal.

## Installation
```shell script
wget https://github.com/matthewpi/privatebin/releases/download/v0.0.1/privatebin -P /usr/bin/
```

## Usage
Currently we only support piping inputs on the Command Line.
```shell script
# Using Echo
echo test | privatebin

# Using Tail
tail -n 20 <FILE> | privatebin

# Using Cat
cat <FILE> | privatebin
```
