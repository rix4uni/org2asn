## org2asn

Extract ASN and IPs using bgp.he.net

## Installation
```
go install github.com/rix4uni/org2asn@latest
```

## Download prebuilt binaries
```
wget https://github.com/rix4uni/org2asn/releases/download/v0.0.3/org2asn-linux-amd64-0.0.3.tgz
tar -xvzf org2asn-linux-amd64-0.0.3.tgz
rm -rf org2asn-linux-amd64-0.0.3.tgz
mv org2asn ~/go/bin/org2asn
```
Or download [binary release](https://github.com/rix4uni/org2asn/releases) for your platform.

## Compile from source
```
git clone --depth 1 github.com/rix4uni/org2asn.git
cd org2asn; go install
```

## Usage
```
Usage of org2asn:
  -json
        Output in JSON format.
  -list string
        File with list of organization names to search ASN info for.
  -org string
        Organization name to search ASN info for.
  -rate-limit int
        Rate limit in seconds between requests.
  -silent
        Silent mode.
  -timeout int
        Timeout for each HTTP request in seconds. (default 10)
  -version
        Print the version of the tool and exit.
```

## Usage Examples
Single ORG
```
▶ echo "dell" | org2asn
```

Multiple ORGs
```
▶ cat orgnames.txt
dell
tesla
yahoo

▶ cat orgnames.txt | org2asn
```

Manually Recursive ORG Search
- Note: make sure you remove manually this will not take too much time, warning you automate this you will lose some results i already tried this and that's why i didn't added this feature.
```
echo "walmart" | org2asn -silent | unew walmart_org2asn.txt

# For advanced use and recursion, remove those not belong to your comapny in walmart_org2asn.txt and run.
cat walmart_org2asn.txt | cut -d"[" -f5 | sed 's/]$//' | org2asn -silent | unew walmart_org2asn.txt

# now remove those not belong to your comapny in walmart_org2asn.txt.
```

## `plain-text` Output
```console
▶ echo "dell" | org2asn

[dell] [AS7977] [ASN] [Dell, Inc.]
[dell] [AS64208] [ASN] [Dell, Inc.]
[dell] [AS59915] [ASN] [Dell Products (Private Unlimited With Share Capital)]
[dell] [AS54701] [ASN] [Dell, Inc.]
[dell] [AS53878] [ASN] [Dell, Inc.]
[dell] [AS46977] [ASN] [Dell, Inc.]
[dell] [AS46507] [ASN] [Dell, Inc.]
[dell] [AS4576] [ASN] [Michael & Susan Dell Foundation]
[dell] [AS38057] [ASN] [Dell (China) Co.,Ltd]
[dell] [AS3615] [ASN] [Dell, Inc.]
[dell] [AS3614] [ASN] [Dell, Inc.]
[dell] [AS3613] [ASN] [Dell, Inc.]
[dell] [AS3612] [ASN] [Dell, Inc.]
[dell] [AS30614] [ASN] [Dell, Inc.]
[dell] [AS23144] [ASN] [Dell, Inc.]
[dell] [AS216418] [ASN] [Filippo Sacco Comis Dell'Oste]
[dell] [AS17187] [ASN] [Dell, Inc.]
[dell] [AS14876] [ASN] [Dell, Inc.]
[dell] [AS12257] [ASN] [Dell, Inc.]
[dell] [65.36.34.0/24] [Route] [Michael & Susan Dell Foundation]
[dell] [63.127.224.0/19] [Route] [DELL COMPUTER (C00824790)]
[dell] [2a12:bec0:430::/44] [Route] [Filippo Sacco Comis Dell'Oste]
[dell] [2620:145:400::/48] [Route] [Dell, Inc.]
[dell] [2607:f2b1:f1e0::/44] [Route] [Dell, Inc.]
[dell] [2607:f2b1:f080::/44] [Route] [Dell, Inc.]
[dell] [2607:f2b1:f000::/42] [Route] [Dell, Inc.]
[dell] [2607:f2b1:e01f::/48] [Route] [Dell, Inc.]
[dell] [2607:f2b1:e019::/48] [Route] [Dell, Inc.]
[dell] [2607:f2b1:e013::/48] [Route] [Dell, Inc.]
[dell] [2607:f2b1:e010::/48] [Route] [Dell, Inc.]
[dell] [2607:f2b1:c050::/44] [Route] [Dell, Inc.]
[dell] [2607:f2b1:c040::/44] [Route] [Dell, Inc.]
[dell] [2607:f2b1::/43] [Route] [Dell, Inc.]
[dell] [2607:f2b1:4001::/48] [Route] [Dell, Inc.]
[dell] [209.222.75.0/24] [Route] [Dell, Inc.]
[dell] [209.222.74.0/24] [Route] [Dell, Inc.]
[dell] [204.75.145.0/24] [Route] [Dell, Inc.]
[dell] [204.75.144.0/24] [Route] [Dell, Inc.]
[dell] [204.197.220.0/24] [Route] [Dell, Inc.]
[dell] [2001:df5:4c00::/48] [Route] [Dell SonicWALL]
[dell] [199.245.235.0/24] [Route] [Dell, Inc.]
[dell] [193.43.18.0/23] [Route] [Agenzia Regionale per la Protezione dell'Ambiente della Sardegna (ARPAS)]
[dell] [193.43.141.0/24] [Route] [Istituto Zooprofilattico Sperimentale dell'Abruzzo e del Molise "G. Caporale"]
[dell] [193.43.117.0/24] [Route] [Agenzia Regionale per la Protezione dell'Ambiente della Sardegna (ARPAS)]
[dell] [192.150.196.0/24] [Route] [Universita' degli Studi dell'Aquila]
[dell] [170.90.8.0/21] [Route] [Dell, Inc.]
[dell] [170.90.79.0/24] [Route] [Dell, Inc.]
[dell] [170.90.74.0/24] [Route] [Dell, Inc.]
[dell] [170.90.64.0/22] [Route] [Dell, Inc.]
[dell] [170.90.4.0/22] [Route] [Dell, Inc.]
[dell] [170.90.16.0/22] [Route] [Dell, Inc.]
[dell] [168.159.228.0/24] [Route] [Dell, Inc.]
[dell] [168.159.224.0/24] [Route] [Dell, Inc.]
[dell] [168.159.218.0/24] [Route] [Dell, Inc.]
[dell] [168.159.215.0/24] [Route] [Dell, Inc.]
[dell] [168.159.212.0/24] [Route] [Dell, Inc.]
[dell] [168.159.209.0/24] [Route] [Dell, Inc.]
[dell] [168.159.175.0/24] [Route] [Dell, Inc.]
[dell] [168.159.174.0/24] [Route] [Dell, Inc.]
[dell] [168.159.173.0/24] [Route] [Dell, Inc.]
[dell] [168.159.172.0/24] [Route] [Dell, Inc.]
[dell] [168.159.171.0/24] [Route] [Dell, Inc.]
[dell] [168.159.170.0/24] [Route] [Dell, Inc.]
[dell] [168.159.169.0/24] [Route] [Dell, Inc.]
[dell] [168.159.168.0/24] [Route] [Dell, Inc.]
[dell] [168.159.167.0/24] [Route] [Dell, Inc.]
[dell] [168.159.166.0/24] [Route] [Dell, Inc.]
[dell] [168.159.165.0/24] [Route] [Dell, Inc.]
[dell] [168.159.162.0/24] [Route] [Dell, Inc.]
[dell] [168.159.161.0/24] [Route] [Dell, Inc.]
[dell] [168.159.160.0/24] [Route] [Dell, Inc.]
[dell] [168.159.158.0/24] [Route] [Dell, Inc.]
[dell] [168.159.157.0/24] [Route] [Dell, Inc.]
[dell] [168.159.156.0/24] [Route] [Dell, Inc.]
[dell] [168.159.155.0/24] [Route] [Dell, Inc.]
[dell] [168.159.154.0/24] [Route] [Dell, Inc.]
[dell] [168.159.153.0/24] [Route] [Dell, Inc.]
[dell] [168.159.152.0/24] [Route] [Dell, Inc.]
[dell] [168.159.151.0/24] [Route] [Dell, Inc.]
[dell] [168.159.150.0/24] [Route] [Dell, Inc.]
[dell] [168.159.149.0/24] [Route] [Dell, Inc.]
[dell] [168.159.148.0/24] [Route] [Dell, Inc.]
[dell] [168.159.147.0/24] [Route] [Dell, Inc.]
[dell] [168.159.146.0/24] [Route] [Dell, Inc.]
[dell] [168.159.145.0/24] [Route] [Dell, Inc.]
[dell] [168.159.144.0/24] [Route] [Dell, Inc.]
[dell] [163.244.72.0/24] [Route] [Dell, Inc.]
[dell] [163.244.64.0/23] [Route] [Dell, Inc.]
[dell] [163.244.62.0/23] [Route] [Dell, Inc.]
[dell] [163.244.60.0/24] [Route] [Dell, Inc.]
[dell] [163.244.57.0/24] [Route] [Dell, Inc.]
[dell] [163.244.56.0/24] [Route] [Dell, Inc.]
[dell] [163.244.40.0/24] [Route] [Dell, Inc.]
[dell] [163.244.38.0/24] [Route] [Dell, Inc.]
[dell] [163.244.37.0/24] [Route] [Dell, Inc.]
[dell] [163.244.246.0/24] [Route] [Dell, Inc.]
[dell] [163.244.24.0/23] [Route] [Dell, Inc.]
[dell] [163.244.185.0/24] [Route] [Dell, Inc.]
[dell] [163.244.12.0/22] [Route] [Dell, Inc.]
[dell] [163.244.116.0/22] [Route] [Dell, Inc.]
[dell] [163.244.0.0/16] [Route] [Dell, Inc.]
[dell] [152.62.60.0/24] [Route] [Dell, Inc.]
[dell] [152.62.47.0/24] [Route] [Dell, Inc.]
[dell] [152.62.45.0/24] [Route] [Dell, Inc.]
[dell] [152.62.44.0/24] [Route] [Dell, Inc.]
[dell] [152.62.185.0/24] [Route] [Dell, Inc.]
[dell] [152.62.184.0/24] [Route] [Dell, Inc.]
[dell] [152.62.176.0/23] [Route] [Dell, Inc.]
[dell] [152.62.168.0/24] [Route] [Dell, Inc.]
[dell] [152.62.166.0/24] [Route] [Dell, Inc.]
[dell] [152.62.155.0/24] [Route] [Dell, Inc.]
[dell] [152.62.154.0/24] [Route] [Dell, Inc.]
[dell] [152.62.118.0/24] [Route] [Dell, Inc.]
[dell] [152.62.115.0/24] [Route] [Dell, Inc.]
[dell] [152.62.113.0/24] [Route] [Dell, Inc.]
[dell] [152.62.109.0/24] [Route] [Dell, Inc.]
[dell] [152.62.108.0/24] [Route] [Dell, Inc.]
[dell] [147.178.9.0/24] [Route] [Dell, Inc.]
[dell] [147.178.8.0/24] [Route] [Dell, Inc.]
[dell] [147.178.7.0/24] [Route] [Dell, Inc.]
[dell] [147.178.6.0/24] [Route] [Dell, Inc.]
[dell] [147.178.5.0/24] [Route] [Dell, Inc.]
[dell] [147.178.4.0/24] [Route] [Dell, Inc.]
[dell] [147.178.252.0/24] [Route] [Dell, Inc.]
[dell] [147.178.240.0/22] [Route] [Dell, Inc.]
[dell] [147.178.2.0/24] [Route] [Dell, Inc.]
[dell] [147.178.192.0/20] [Route] [Dell, Inc.]
[dell] [143.166.208.0/21] [Route] [Dell, Inc.]
[dell] [143.166.200.0/21] [Route] [Dell, Inc.]
[dell] [143.166.15.0/24] [Route] [Dell, Inc.]
[dell] [143.166.0.0/16] [Route] [Dell, Inc.]
[dell] [137.69.127.0/24] [Route] [Dell, Inc.]
[dell] [137.69.126.0/24] [Route] [Dell, Inc.]
[dell] [137.69.120.0/24] [Route] [Dell, Inc.]
[dell] [137.69.118.0/24] [Route] [Dell, Inc.]
[dell] [137.69.117.0/24] [Route] [Dell, Inc.]
[dell] [137.69.113.0/24] [Route] [Dell, Inc.]
[dell] [137.69.112.0/24] [Route] [Dell, Inc.]
[dell] [132.237.196.0/24] [Route] [Dell, Inc.]
[dell] [132.237.194.0/24] [Route] [Dell, Inc.]
[dell] [132.237.192.0/24] [Route] [Dell, Inc.]
[dell] [132.237.186.0/24] [Route] [Dell, Inc.]
[dell] [132.237.185.0/24] [Route] [Dell, Inc.]
[dell] [132.237.184.0/24] [Route] [Dell, Inc.]
[dell] [132.237.181.0/24] [Route] [Dell, Inc.]
[dell] [132.237.180.0/24] [Route] [Dell, Inc.]
[dell] [132.237.18.0/23] [Route] [Dell, Inc.]
[dell] [132.237.178.0/24] [Route] [Dell, Inc.]
[dell] [132.237.177.0/24] [Route] [Dell, Inc.]
[dell] [132.237.173.0/24] [Route] [Dell, Inc.]
[dell] [132.237.171.0/24] [Route] [Dell, Inc.]
[dell] [132.237.170.0/24] [Route] [Dell, Inc.]
[dell] [132.237.169.0/24] [Route] [Dell, Inc.]
[dell] [132.237.167.0/24] [Route] [Dell, Inc.]
[dell] [132.237.166.0/24] [Route] [Dell, Inc.]
[dell] [132.237.164.0/24] [Route] [Dell, Inc.]
[dell] [132.237.161.0/24] [Route] [Dell, Inc.]
[dell] [132.237.16.0/23] [Route] [Dell, Inc.]
[dell] [132.237.16.0/22] [Route] [Dell, Inc.]
[dell] [132.237.156.0/24] [Route] [Dell, Inc.]
[dell] [132.237.155.0/24] [Route] [Dell, Inc.]
[dell] [132.237.153.0/24] [Route] [Dell, Inc.]
[dell] [132.237.151.0/24] [Route] [Dell, Inc.]
[dell] [132.237.146.0/24] [Route] [Dell, Inc.]
[dell] [132.237.141.0/24] [Route] [Dell, Inc.]
[dell] [132.237.139.0/24] [Route] [Dell, Inc.]
[dell] [132.237.136.0/24] [Route] [Dell, Inc.]
[dell] [132.237.135.0/24] [Route] [Dell, Inc.]
[dell] [132.237.134.0/24] [Route] [Dell, Inc.]
[dell] [132.237.132.0/24] [Route] [Dell, Inc.]
[dell] [132.237.130.0/23] [Route] [Dell, Inc.]
[dell] [132.237.128.0/23] [Route] [Dell, Inc.]
[dell] [128.221.239.0/24] [Route] [Dell, Inc.]
[dell] [128.221.238.0/24] [Route] [Dell, Inc.]
[dell] [128.221.237.0/24] [Route] [Dell, Inc.]
[dell] [128.221.236.0/24] [Route] [Dell, Inc.]
[dell] [128.221.235.0/24] [Route] [Dell, Inc.]
[dell] [128.221.234.0/24] [Route] [Dell, Inc.]
[dell] [128.221.233.0/24] [Route] [Dell, Inc.]
[dell] [128.221.231.0/24] [Route] [Dell, Inc.]
[dell] [128.221.230.0/24] [Route] [Dell, Inc.]
[dell] [128.221.228.0/24] [Route] [Dell, Inc.]
[dell] [128.221.227.0/24] [Route] [Dell, Inc.]
[dell] [128.221.226.0/24] [Route] [Dell, Inc.]
[dell] [128.221.225.0/24] [Route] [Dell, Inc.]
[dell] [128.221.224.0/24] [Route] [Dell, Inc.]
[dell] [128.221.223.0/24] [Route] [Dell, Inc.]
[dell] [128.221.222.0/24] [Route] [Dell, Inc.]
[dell] [128.221.221.0/24] [Route] [Dell, Inc.]
[dell] [128.221.220.0/24] [Route] [Dell, Inc.]
[dell] [128.221.219.0/24] [Route] [Dell, Inc.]
[dell] [128.221.218.0/24] [Route] [Dell, Inc.]
[dell] [128.221.217.0/24] [Route] [Dell, Inc.]
[dell] [128.221.215.0/24] [Route] [Dell, Inc.]
[dell] [128.221.214.0/24] [Route] [Dell, Inc.]
[dell] [128.221.213.0/24] [Route] [Dell, Inc.]
[dell] [128.221.211.0/24] [Route] [Dell, Inc.]
[dell] [128.221.210.0/24] [Route] [Dell, Inc.]
[dell] [128.221.209.0/24] [Route] [Dell, Inc.]
[dell] [128.221.208.0/24] [Route] [Dell, Inc.]
[dell] [128.221.204.0/24] [Route] [Dell, Inc.]
[dell] [128.221.192.0/24] [Route] [Dell, Inc.]
[dell] [103.19.169.0/24] [Route] [Dell SonicWALL]
[dell] [103.19.168.0/24] [Route] [Dell SonicWALL]
[dell] [103.13.13.0/24] [Route] [Dell (China) Co.,Ltd]
[dell] [103.13.12.0/24] [Route] [Dell (China) Co.,Ltd]
```
