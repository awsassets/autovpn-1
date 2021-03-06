# Autovpn

Manages Openvpn connections, it start automatically once Internet connection has been established. In addition, it will choose a random openvpn configuration file (different vpn server) every time is restarted.


## Configuration

The following commands will create a configuration directory in `$HOME/.autovpn` and compile the tool:

```
make init
make build
```

Finally, we just need configure our Openvpn credentials and vpn configuration files:

1- Create a file named `openvpn.creds` with your username and password and copy it to `$HOME/.autovpn`. Openvpn credentials files contain two lines, one for the username and the second one for the password, e.g:

```
myusername
mysuperpassword
```

2- Copy your VPNs files into `$HOME/.autovpn/confs`, `autovpn` will choose a random one to launch a new vpn connection every time it restarts

### .autovpn structure

```
.
├── confs
│   ├── de.protonvpn.com.udp.ovpn
│   ├── es.protonvpn.com.udp.ovpn
│   ├── fr.protonvpn.com.udp.ovpn
│   └── pt.protonvpn.com.udp.ovpn
└── openvpn.creds

1 directory, 5 files
```

## Execution

**Note:** VPN tunnels require root permissions

```
make start

//or

cd autobuild && sudo ./autovpn
```



