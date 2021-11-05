# Docker for Mac on M1Pro laptop

Quick and dirty test on my 10core M1Pro Macbook Pro.



```
docker-compose up


// ... lots of stuff

mysql-benchmark-bench-1  | starting benchmark for test:example@(mysql-rosetta)/testdb
mysql-benchmark-bench-1  | starting setup for test:example@(mysql-rosetta)/testdb
mysql-benchmark-bench-1  | finished inserting 100000 rows from 100 goroutines for test:example@(mysql-rosetta)/testdb in 14.741371381s
mysql-benchmark-bench-1  | starting benchmark for test:example@(mysql-native)/testdb
mysql-benchmark-bench-1  | starting setup for test:example@(mysql-native)/testdb
mysql-benchmark-bench-1  | finished inserting 100000 rows from 100 goroutines for test:example@(mysql-native)/testdb in 10.03471088s
mysql-benchmark-bench-1 exited with code 0
```