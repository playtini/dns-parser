# Tool for parsing dns records

## This is a fork of https://github.com/majek/goplayground/tree/master/resolve

### Changes

- removed '6' option

- added 'dns-type' options (dnsTypeA = 1, dnsTypeNS = 2, dnsTypeMX = 15, dnsTypeTXT = 16, dnsTypeAAAA  = 28), default - dnsTypeA

### Usage example 

`go build; ./dns-parser -v --dns-type=15 < input.txt > res.txt`

input.txt - file with domain per line