# parserdor

Parser domain,subdomain,url, ip or cidr by stdin

## Usage Example

```
▶ cat scope.txt | parserdor
[url] https://hackerone.com
[subdomain] ext.hackerone.com
[domain] hackerone.com
[ip] 192.168.1.1
[cidr] 192.168.1.0/24
```

## Install

```
go install github.com/phor3nsic/parserdor@latest
```