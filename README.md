# snicat

snicat is a TLS aware netcat.

It is specifically designed for connecting to services using SNI multiplexing (TLS virtual hosts). That is multiple services using a single port behind a reverse proxy.

```
./sc <hostname> <port>

Usage of ./sc:
  -b string
        Shorthand for -bind
  -bind string
        Tunnel connection to a local unencrypted port (e.g. localhost:3000, 3000)
  -insecure
        Enable verification of server SSL certificate
  -k    Shorthand for -insecure
  -servername string
        Server Name Indication (SNI) to provide to the server (e.g. ssl.example.com)
  -sni string
        Shorthand for -servername
```

# Install

Download the latest release from Github for your appropriate operating system.

For Linux and Mac the following commands should download snicat for you:

```
wget "https://github.com/CTFd/snicat/releases/latest/download/sc_`uname`_`uname -m`" -O sc
chmod +x sc
```

For Mac, you may need to remove the quarantine attribute by running:

```
xattr -d com.apple.quarantine sc
```

For Windows, you can download the x64 exe [here](https://github.com/CTFd/snicat/releases/latest/download/sc_Windows_x86_64.exe) and then run it from the command prompt.

# Examples

## Basic Usage

```
./sc ssl.example.com
```

## Tunnel Remote Port Locally

```
./sc -bind 3000 ssl.example.com
```

Then you can subsequently connect to the remote host through a local proxy. For example:

```
nc localhost 3000
```

# Why would I need this?

## TLS/SSL aware TCP servers

Sometimes, perhaps for CTFs, you may encounter services behind an SSL-enabled reverse proxy. Because the reverse proxy is encrypting the connection, it may be difficult for some tools (i.e. netcat) to connect and interact with the service.

For example, [slt](https://github.com/inconshreveable/slt) is a tool that provides an SNI proxy for arbitrary TCP services.

snicat can establish the encrypted connection and expose a local unencrypted connection or you can use snicat to do all your communication.

## What if I dont want to install a new tool!

You don't have to! snicat is mostly an easier way to do something that you can probably already do.

**OpenSSL**

```
openssl s_client -connect ssl.example.com:443 -servername ssl.example.com -quiet
```

**Python**

We provide a [version of snicat written in Python](python) with no external dependencies that is easy to review and install.

# Acknowledgements

The majority of the functionality provided by snicat is provided by [sclient](https://github.com/therootcompany/sclient). In the future we may choose to no longer rely on it but we thank it for providing initial code and inspiration.