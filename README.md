# Investing Helper

This program help us to calc weight for stock on stock exchange.

Call:
```sh
run main.go [price] [input file path] [output file path]
```
Example:
```sh
run main.go 2000.00 input.csv output.csv
```

The input.csv file example:
```
B3SA3,12
BTOW3,10
CSAN3,10
CYRE3,8
EZTC3,8
MGLU3,10
PETR4,10
VALE3,12
VVAR3,10
WEGE3,10
```

The output.csv file example:
```
stock,weight,value,quantity,total
B3SA3,12,62.34,3,187.02
BTOW3,10,112.81,1,112.81
CSAN3,10,83.81,2,167.62
CYRE3,8,24.43,6,146.58
EZTC3,8,40.10,3,120.30
MGLU3,10,87.35,2,174.70
PETR4,10,22.73,8,181.84
VALE3,12,62.95,3,188.85
VVAR3,10,19.55,10,195.50
WEGE3,10,67.22,2,134.44
-,-,-,-,1609.66
```