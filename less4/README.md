Run

```bash
$ go run ./less4/formatter/formatter.go 
2019/05/25 03:05:27 Initializing logging reporter
Running Formatter at :8081
2019/05/25 03:05:34 Reporting span 125052673eb886b5:52cde3bcda1b607e:6e66fd82b548aae2:1

```

```bash
$  go run ./less4/publisher/publisher.go
2019/05/25 04:03:29 Initializing logging reporter
Running Publisher at :8082
Hi sanjid
2019/05/25 04:03:37 Reporting span 2846e480c4bf519d:4a27013e070d13b8:530a362281e05dde:1

```


```bash
$  go run ./less4/client/hello.go sanjid Hi
2019/05/25 03:05:34 Initializing logging reporter
2019/05/25 03:05:34 Reporting span 125052673eb886b5:6e66fd82b548aae2:125052673eb886b5:1
2019/05/25 03:05:34 Reporting span 125052673eb886b5:ccad60db1fa39c8:125052673eb886b5:1
2019/05/25 03:05:34 Reporting span 125052673eb886b5:125052673eb886b5:0:1
 
```
