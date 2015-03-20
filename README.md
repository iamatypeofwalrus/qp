# Query Parse
Parse URL query params from the command line.

## Usage
```sh
# Prints
# xml	2
echo /review/list?format=xml&v=2 | qp -param 'format,v'

# Verbose prints a tab list of
# 	url
#		query string
#		values pulled out of query string
#
# /review/list	format=xml&v=2	xml
echo /review/list?format=xml&v=2 | qp -param 'format' -verbose
```

## Notes
When multiple params are specified and at least one is found any empty param will be signified as ".".

This allows to guarantee the position of the param regardless of the presence of other parameters based on the order entered in the -param option.
