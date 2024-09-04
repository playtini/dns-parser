# Tool for parsing dns records

## This is a fork of https://github.com/majek/goplayground/tree/master/resolve

### Changes

- removed '6' option

- added 'dns-type' options (default - dnsTypeA)
  - dnsTypeA = 1
  - dnsTypeNS = 2
  - dnsTypeMX = 15
  - dnsTypeTXT = 16
  - dnsTypeAAAA = 28

### Usage example 

`go build; ./dns-parser -v --dns-type=1 < ./input.txt > ./output.txt`

input.txt - file with domain per line