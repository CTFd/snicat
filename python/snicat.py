#!/usr/bin/env python3
from telnetlib import Telnet
import argparse
import ssl

parser = argparse.ArgumentParser(description="Connect to services behind an SNI proxy.")

parser.add_argument("hostname", type=str, help="hostname to connect to")
parser.add_argument(
    "port", type=int, default=443, nargs="?", help="port to connect to",
)
parser.add_argument(
    "--insecure",
    "-k",
    action="store_true",
    help="Disable verification of server SSL certificate",
)
parser.add_argument(
    "--servername",
    "-sni",
    type=str,
    default=None,
    nargs="?",
    help="Server Name Indication (SNI) to provide to the server (e.g. ssl.example.com)",
)

args = parser.parse_args()

context = ssl.create_default_context()
if args.insecure:
    context.check_hostname = False
    context.verify_mode = ssl.CERT_NONE

tn = Telnet(args.hostname, args.port)

servername = args.servername
if servername is None:
    servername = args.hostname

with context.wrap_socket(tn.sock, server_hostname=servername) as secure_sock:
    tn.sock = secure_sock
    tn.interact()
