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
BBDC4,10
CSAN3,10
CYRE3,10
GGBR4,10
MGLU3,7
PETR4,12
VALE3,12
VVAR3,7
WEGE3,10
```

The output.csv file example:
```
stock,weight,value,quantity,total
B3SA3,12,57.82,2,115.64
BBDC4,10,21.95,5,109.75
CSAN3,10,76.70,1,76.70
CYRE3,10,24.30,5,121.50
GGBR4,10,19.38,6,116.28
MGLU3,7,88.50,1,88.50
PETR4,12,22.88,6,137.28
VALE3,12,59.32,2,118.64
VVAR3,7,19.19,4,76.76
WEGE3,10,64.43,1,64.43
-,-,-,1025.48,1264.29
```