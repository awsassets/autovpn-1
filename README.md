# Autovpn

Manages Openvpn connections, it start automatically once Internet connection has been established. In addition, it will choose a random openvpn configuration file (different vpn server) every time is restarted.


## Configuration

The following commands will create a configuration directory in `$HOME/.autovpn` and compile the tool:

```
make init
make build
```

Finally, we just need configure our Openvpn credentials and vpn configuration files:

1- Create a file named `creds.openvpn` with your username and password and copy it to `$HOME/.autovpn`. Openvpn credentials files contain two lines, one for the username and the second one for the password, e.g:

```
myusername
mysuperpassword
```

2- Copy your VPNs files into `$HOME/.autovpn/confs`, `autovpn` will choose a random one to launch a new vpn connection every time it restarts

