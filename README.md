# org2asn
Extract ASN in bgp.he.net

# Install
```
git clone https://github.com/rix4uni/org2asn.git
cd org2asn
pip3 install -r requirements.txt
```

# Usage
```
echo "dell" | python3 org2asn.py
```

```
cat names.txt | python3 org2asn.py

## names.txt file contains
dell
tesla
yahoo
```

# Output
<details>
<summary><b>dell</b></summary>

```json
{
  "org": "dell",
  "num_ASN": 18,
  "num_Route": 114,
  "details": [
    {
      "Result": "AS7977",
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
      "Result": "AS46507",
      "Type": "ASN",
      "Description": "Dell, Inc."
    },
    {
      "Result": "AS4576",
      "Type": "ASN",
      "Description": "Michael & Susan Dell Foundation"
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
      "Description": "Michael & Susan Dell Foundation"
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
      "Result": "2607:f2b1:e01f::/48",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "2607:f2b1:e010::/48",
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
      "Result": "2001:df5:4c00::/48",
      "Type": "Route",
      "Description": "Dell SonicWALL"
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
      "Result": "192.150.194.0/23",
      "Type": "Route",
      "Description": "Universita' degli Studi dell'Aquila"
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
      "Result": "168.159.172.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.169.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.168.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.163.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.160.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "168.159.155.0/24",
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
      "Result": "168.159.145.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.9.0/24",
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
      "Result": "163.244.22.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.186.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.185.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.180.0/24",
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
      "Result": "163.244.10.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "163.244.0.0/16",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "147.178.9.0/24",
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
      "Result": "143.166.0.0/16",
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
      "Result": "132.237.176.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.174.0/24",
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
      "Result": "132.237.158.0/24",
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
      "Result": "132.237.148.0/23",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.146.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "132.237.141.0/24",
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
      "Result": "128.221.234.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.232.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.231.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.228.0/24",
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
      "Result": "128.221.219.0/24",
      "Type": "Route",
      "Description": "Dell, Inc."
    },
    {
      "Result": "128.221.218.0/24",
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
</details>
