# GameMaster Client (gmc)

## Build

1. Build
        go build
2. Run
        ./gmc -ip=<gm_ip> -port=<port>
If flags are not provided, will default to 127.0.0.1:50000

3. Example

First, make sure gmd is running. Without it, client will have nowhere to connect to.

Lets start the client:
        ./gmc

If we have a game at contract 0x7B9d950cC1ecD94eD0cF3916989B0ac56C70AB24,
we connect it like so:
        gmclerk> connect  0x7B9d950cC1ecD94eD0cF3916989B0ac56C70AB24

The relevant logging should be triggered on the gmd end. If everything went well,
gmc will return the prompt to you. If something went bad, it will print the error and return the promt.  Restart gmc if you are getting this error:
        2018/06/13 04:28:47 connection is shut down

If we now want to disconnect the game, here is how we do it:
        gmclerk> disconnect  0x7B9d950cC1ecD94eD0cF3916989B0ac56C70AB24

We can also connect all games in a file:
        gmclerk> connect -f <filename>