# snicat (Python)

We provide a less full featured Python implemention of snicat.

This version of snicat is intended for those who would like to run code that they can easily review or would otherwise prefer Python code over the normal Golang code.

The Python snicat does not have the tunneling feature provided by the main snicat. PRs are welcome to implement this functionality.

```
❯ ./snicat.py -h
usage: snicat.py [-h] [--insecure] [--servername [SERVERNAME]] hostname [port]

Connect to services behind an SNI proxy.

positional arguments:
  hostname              hostname to connect to
  port                  port to connect to

optional arguments:
  -h, --help            show this help message and exit
  --insecure, -k        Disable verification of server SSL certificate
  --servername [SERVERNAME], -sni [SERVERNAME]
                        Server Name Indication (SNI) to provide to the server (e.g. ssl.example.com)
```

Some examples of invocation:

```
./snicat.py ssl.example.com
```

```
./snicat.py ssl.example.com 12345
```

```
python3 snicat.py ssl.example.com 12345 --insecure
```

```
python3 snicat.py ssl.example.com 12345 --servername echo.service --insecure
```