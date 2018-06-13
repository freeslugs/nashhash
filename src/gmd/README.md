# Game Master Daemon

## Build

    go build

## Run

    gmd -ip=<ip_addr> -port=<port_number> -key=<owner_private_key>

We can omit -ip and -port flags to run gmd on 127.0.0.1:50000
Additionally, port provided must be a private port. 

gmd is not daemonized yet. To run it in the background you will have to explicitly do so.

## Example

Lets run the gmd in the background

    nohup ./gmd -key=76a23cff887b294bb60ccde7ad1eb800f0f6ede70d33b154a53eadb20681a4e3 &>log.txt &

This above command will report the PID of the process.
Use this PID to kill the process if needed.

If we need to kill it
    kill -9 <PID>
