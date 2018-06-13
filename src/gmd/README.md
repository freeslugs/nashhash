# Game Master Daemon

## Build
    go build

## Run
    gmd -ip=<ip_addr> -port=<port_number> -key=<owner_private_key>
We can omit -ip and -port flags to run gmd on 127.0.0.1:50000
Additionally, port provided must be a private port. 