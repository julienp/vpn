# ExpressVPNController

HTTP API for the ExpressVPN [Linux App](https://www.expressvpn.com/support/vpn-setup/app-for-linux/#).

The goal is to have this running on a RaspberryPi with ExpressVPN that acts as
a router. This API should allow an app to control the VPN connection, chosing
countries, etc.

## Build
 
env GOOS=linux GOARCH=arm GOARM=5 go build -o pivpn main.go