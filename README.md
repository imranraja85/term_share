# Idea: I want to be able to spin up a server that broadcasts my tmux session.
Anyone can connect and watch what i do (only read me initially)

## first i start the server and listen on a particular ip and port
./term_share -ip <the-ip-and-port-to-bind-to> -isHost true


## then a client would connect to the ip
./term_share -ip <the-ip-and-port-to-listen-on>


### Current State
I have a listener running on 127.0.0.1:3009 inside a goroutine

I have a vim session running in a separate process with its STDOUT, STDIN and STDERR set to the current os STD

What do I want to try next?
Sending the STDOUT from the externel vim process over a channel!

### Redirect output of externel process to stdout: http://stackoverflow.com/questions/8875038/redirect-stdout-pipe-of-child-process-in-golang
