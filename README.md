# log-go-grep

A lightweight CLI tool for searching patterns in log files. 
Mainly to teach myself building GO CLI tools

## Installation

```bash
go build
```

## Usage

```bash
logGoGrep -s "pattern" [file]
```

### Flags

- `-s, --search` - Search pattern (required)
- `-f, --file` - Log file path (optional, can use positional argument)
- `-c, --count` - Display count of matching lines

### Examples

```bash
# Basic search
logGoGrep -s "error" system.log

# With count
logGoGrep -s "error" -c system.log

# Using file flag
logGoGrep -s "warning" -f system.log
```

The file can be specified as a flag (`-f`) or as the last positional argument.
