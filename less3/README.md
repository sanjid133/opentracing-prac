Run

```bash
$ go run ./less3/formatter/formatter.go 
2019/05/25 03:05:27 Initializing logging reporter
Running Formatter at :8081
2019/05/25 03:05:34 Reporting span 125052673eb886b5:52cde3bcda1b607e:6e66fd82b548aae2:1

```

```bash
$  go run ./less3/publisher/publisher.go
2019/05/25 03:05:27 Initializing logging reporter
Running Formatter at :8081
2019/05/25 03:05:34 Reporting span 125052673eb886b5:52cde3bcda1b607e:6e66fd82b548aae2:1
 
```


```bash
$  go run ./less3/publisher/publisher.go
2019/05/25 03:05:34 Initializing logging reporter
2019/05/25 03:05:34 Reporting span 125052673eb886b5:6e66fd82b548aae2:125052673eb886b5:1
2019/05/25 03:05:34 Reporting span 125052673eb886b5:ccad60db1fa39c8:125052673eb886b5:1
2019/05/25 03:05:34 Reporting span 125052673eb886b5:125052673eb886b5:0:1
 
```
