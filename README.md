## org2asn

Extract ASN and IPs in bgp.he.net

## Installation
```
go install github.com/rix4uni/org2asn@latest
```

## Download prebuilt binaries
```
wget https://github.com/rix4uni/org2asn/releases/download/v0.0.1/org2asn-linux-amd64-0.0.1.tgz
tar -xvzf org2asn-linux-amd64-0.0.1.tgz
rm -rf org2asn-linux-amd64-0.0.1.tgz
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
echo "dell" | python3 org2asn
```

```
cat orgnames.txt | python3 org2asn

## orgnames.txt file contains
dell
tesla
yahoo
```

## Output
```json
{
  "org": "dell",
  "num_ASN": 21,
  "num_Route": 183,
  "details": [
    {
      "Result": "AS7977",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS64208",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS59915",
      "Type": "ASN",
      "Description": "Dell Products (Private Unlimited With Share Capital)"
    },
    {
      "Result": "AS54701",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS53878",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS46977",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS46507",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS4576",
      "Type": "ASN",
      "Description": "Michael \u0026 Susan Dell Foundation"
    },
    {
      "Result": "AS38057",
      "Type": "ASN",
      "Description": "Dell (China) Co.,Ltd"
    },
    {
      "Result": "AS3615",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS3614",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS3613",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS3612",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS30614",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS23144",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS216418",
      "Type": "ASN",
      "Description": "Filippo Sacco Comis Dell'Oste"
    },
    {
      "Result": "AS17919",
      "Type": "ASN",
      "Description": "Dell AP-Online"
    },
    {
      "Result": "AS17187",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS14876",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS132711",
      "Type": "ASN",
      "Description": "Dell SonicWALL"
    },
    {
      "Result": "AS12257",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "65.36.34.0/24",
      "Type": "Route",
      "Description": "Michael \u0026 Susan Dell Foundation"
    },
    {
      "Result": "63.127.224.0/19",
      "Type": "Route",
      "Description": "DELL COMPUTER (C00824790)"
    },
    {
      "Result": "2a12:bec0:430::/44",
      "Type": "Route",
      "Description": "Filippo Sacco Comis Dell'Oste"
    },
    {
      "Result": "2620:145:800::/38",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2620:145:400::/48",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2607:f2b1:f1e0::/44",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2607:f2b1:f080::/44",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2607:f2b1:f000::/42",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2607:f2b1:e01f::/48",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2607:f2b1:e019::/48",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2607:f2b1:e013::/48",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2607:f2b1:e010::/48",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2607:f2b1:c050::/44",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2607:f2b1:c040::/44",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2607:f2b1::/43",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2607:f2b1:71::/48",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2607:f2b1:70::/48",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2607:f2b1:4001::/48",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "209.222.75.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "209.222.74.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "204.75.145.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "204.75.144.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2001:df5:4c00::/48",
      "Type": "Route",
      "Description": "Dell SonicWALL"
    },
    {
      "Result": "199.245.235.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "193.43.18.0/23",
      "Type": "Route",
      "Description": "Agenzia Regionale per la Protezione dell'Ambiente della Sardegna (ARPAS)"
    },
    {
      "Result": "193.43.141.0/24",
      "Type": "Route",
      "Description": "Istituto Zooprofilattico Sperimentale dell'Abruzzo e del Molise \"G. Caporale\""
    },
    {
      "Result": "193.43.117.0/24",
      "Type": "Route",
      "Description": "Agenzia Regionale per la Protezione dell'Ambiente della Sardegna (ARPAS)"
    },
    {
      "Result": "192.150.196.0/24",
      "Type": "Route",
      "Description": "Universita' degli Studi dell'Aquila"
    },
    {
      "Result": "170.90.8.0/21",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "170.90.79.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "170.90.74.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "170.90.64.0/22",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "170.90.64.0/18",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "170.90.4.0/22",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "170.90.32.0/19",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "170.90.16.0/22",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "170.90.128.0/17",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.228.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.224.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.218.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.215.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.212.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.209.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.175.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.174.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.173.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.172.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.171.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.170.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.169.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.167.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.166.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.165.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.162.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.161.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.160.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.158.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.157.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.156.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.155.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.154.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.153.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.152.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.151.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.150.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.149.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.148.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.147.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.146.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.145.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.144.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.72.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.64.0/23",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.62.0/23",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.60.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.57.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.56.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.40.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.38.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.37.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.246.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.24.0/23",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.185.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.12.0/22",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.116.0/22",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.0.0/16",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "152.62.47.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "152.62.45.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "152.62.44.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "152.62.185.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "152.62.184.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "152.62.176.0/23",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "152.62.168.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "152.62.166.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "152.62.155.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "152.62.154.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "152.62.113.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "152.62.109.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "152.62.108.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "147.178.9.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "147.178.8.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "147.178.6.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "147.178.4.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "147.178.252.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "147.178.240.0/22",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "147.178.2.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "147.178.192.0/20",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "143.166.208.0/21",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "143.166.200.0/21",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "143.166.15.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "143.166.0.0/16",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "137.69.127.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "137.69.126.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "137.69.120.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "137.69.118.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "137.69.117.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "137.69.113.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "137.69.112.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.196.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.194.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.192.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.186.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.185.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.184.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.181.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.180.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.18.0/23",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.178.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.177.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.173.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.171.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.170.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.169.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.167.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.166.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.164.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.163.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.161.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.16.0/23",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.16.0/22",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.156.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.155.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.153.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.151.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.146.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.139.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.136.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.135.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.134.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.132.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.130.0/23",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.128.0/23",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.239.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.238.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.237.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.236.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.235.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.234.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.233.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.231.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.230.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.227.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.226.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.225.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.224.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.223.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.222.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.221.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.220.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.218.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.217.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.215.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.214.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.213.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.211.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.209.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.208.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.204.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.192.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "103.19.169.0/24",
      "Type": "Route",
      "Description": "Dell SonicWALL"
    },
    {
      "Result": "103.19.168.0/24",
      "Type": "Route",
      "Description": "Dell SonicWALL"
    },
    {
      "Result": "103.13.13.0/24",
      "Type": "Route",
      "Description": "Dell (China) Co.,Ltd"
    },
    {
      "Result": "103.13.12.0/24",
      "Type": "Route",
      "Description": "Dell (China) Co.,Ltd"
    }
  ]
}
```
