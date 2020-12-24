# Net-Tools

### A collection of few basic tools to help in common networking applications

## Subnet Calculator

This tool allows the user to type in an ip address and get all the attributes about the subnet that the device is connected in. Currently only /25 and above subnets are implemented. Higher subnets to be added soon. The following attributes of the network will be returned.

1. Network ID
2. Broadcast IP
3. First IP
4. Last IP
5. Next Network ID
6. Usable IP addresses
7. Number of hosts in the network
8. CIDR of the subnet

#### To use the subnet calculator

```
cd Subnet
go run main.go

Type in an IPv4 address of the form a.b.c.d/x where x belongs to /25 or above subnet

```

## Scanner

This tool allows the user to view all the devices that are connected to the same network as that of the user. It is similar to the nmap command in Linux based OS.

#### To use the host scanner

```
cd Scanner
go run main.go
```

## HTTP Client

A simple command line tool that lets you download resources over HTTP. The tool takes the website address and path of the resource on the server and downloads the resource on the host machine.

#### To use the client program

```
cd Scanner
go run main.go sitename path resource
```

## Torrent Parser

A tool that lets you parse torrent files to a readable format. Shows the list of trackers, files and other meta-data realted to the torrent
