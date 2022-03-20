# OpenVPN HTTP API

Simple extension for OpenVPN providing HTTP interface for `.ovpn` file generation.

## Run
1. Copy config sample `cp config.yml.sample config.yml`. Configure it:
    * Set CA private key passphrase with `ca_private_key_pass` parameter;
    * Set port HTTP API will run on with `server.port` parameter.
1. Follow [this manual](https://github.com/kylemanna/docker-openvpn/blob/master/docs/docker-compose.md) for initial OpenVPN config
1. Run all the stuff with `docker-compose up --build -d`

## Use
1. Try `curl -X POST localhost:<your_port>/ovpn-config?clientId=new_client&password=new_password`
