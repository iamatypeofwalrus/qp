# Query Parse
Parse Url query params from the command line

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
