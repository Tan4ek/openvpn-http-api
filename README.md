# OpenVPN HTTP API

Simple extension for OpenVPN providing HTTP interface for `.ovpn` file generation.

The management API has no authentication and can issue or retrieve VPN client
configuration. Compose publishes it only on `127.0.0.1:8090`; the VPN UDP port
1194 remains public. Access management locally or through an authenticated
tunnel. Other trusted containers can use port 8090 over a shared Docker network.
Do not expose the management port publicly without an authenticated proxy.

Existing containers need to be recreated from the updated Compose file before
this restriction takes effect. The legacy Go/OpenVPN images need a separate
runtime upgrade; changing the port binding does not update them.

## Run
1. Copy config sample `cp config.yml.sample config.yml`. Configure it:
    * Set CA private key passphrase with `ca_private_key_pass` parameter;
    * Set port HTTP API will run on with `server.port` parameter.
1. Follow [this manual](https://github.com/kylemanna/docker-openvpn/blob/master/docs/docker-compose.md) for initial OpenVPN config
1. Run all the stuff with `docker-compose up --build -d`

## Use
1. Try `curl -X POST localhost:<your_port>/ovpn-config?clientId=new_client&password=new_password`
