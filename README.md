# Golang File Formatter (GOFF)

This is a very quick util I started building to help convert CSV to JSON. Right now it only supports CSV > JSON with minimal error recovery.

## Usage

Assuming the following CSV file named `data.csv`:

```csv
ID,Name,Salary
1,James Murphy,32000.99
2,Jane Doe,44000
3,John Doe,12345
```

... the output of `goff data.csv` would result in a `data.json` file with the following contents:

```json
[
	{
		"ID": "1",
		"Name": "James Murphy",
		"Salary": "32000.99"
	},
	{
		"ID": "2",
		"Name": "Jane Doe",
		"Salary": "44000"
	},
	{
		"ID": "3",
		"Name": "John Doe",
		"Salary": "12345"
	}
]
```

Usage is `goff <input file path> <output file path>`. Example commands:

```bash 
# Read data.csv and output myjsondata.json:
goff data.csv myjsondata.json
# Read data.csv and default second param to <inputFileName>.json. In this case data.json:
goff data.csv
```

## Installation

Make sure you have [golang  installed](https://golang.org/) and your gopath set, then run `go get`:

```bash
go get github.com/DeanPDX/goff/cmd/goff
```

If everything goes as planned, you should be able to run `goff` and see the usage notes:

```bash
v0.1
usage:

goff <input file path> <output file path>

the second argument (output file path) is optional
and will default to the name of your input file with
the appropriate extension.
```

## Limitations

- We only support CSV > JSON. Consider adding other file types.
- Data types are always strings. Consider being more clever about data types for formats that support more complex types (like JSON).
- Malformed file error handling could be better. See [this question in StackOverflow](https://stackoverflow.com/questions/30633115/golang-file-reading-only-reading-last-line). We have a similar problem in the event that incorrect newlines are used. Consider using [transform.Transformer](https://godoc.org/golang.org/x/text/transform#Transformer) to clean up malformed files.

## Future Updates

- Convert read/write to buffered IO. This will allow easier unit testing.
- Get some test coverage.