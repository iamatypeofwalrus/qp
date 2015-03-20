# Query Parse
Parse URL query params from the command line.

## Usage
```
Usage: qp [options] <comma delimited list of param keys>
  -null="NULL": Value to print when a param is not present.
  -v=false: Print url and query string in addition to values matching params

Example: qp -v -null 'derp' 'format, sort'
```

## Examples
```sh
echo /review/list?format=xml&v=2 | qp 'format,v'
# Prints:
# xml	2

# Verbose prints a tab list of
# 	$1    -> url
#		$2    -> query string
#		$3... -> values pulled out of query string
echo /review/list?format=xml&v=2 | qp -param 'format, derp' -verbose
# Prints:
# /review/list	format=xml&v=2	xml NULL
```

## Notes
When multiple params are specified and at least one value in a given line is present then any null value will print the null string. By default that is NULL.

If all parameters are NULL or there is any error parsing the URL there will be no output for that line.

This guarantees the position of the parameter value in the output. Output order is based on the input order.
