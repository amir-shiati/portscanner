# Portscanner
A fast and simple port scanner created in golang. for times when you are not able to use nmap!

# Usage
` Port scanner [-h|--help] -i|--ip "<value>" [-p|--port "<value>"]
                    [-l|--protocol "<value>"] [-c|--common "<value>"]
                    [-r|--range "<value>"] [-d|--duration "<value>"]
`
### clone the repo.
`$ git clone https://github.com/amir-shiati/portscanner.git`
### Change directory.
`$ cd portscanner/`
### Give the binary file permission to execute ***if needed***.
`$ chmod -R 775 ./port_scanner`
### Simply run the binary file
`$ ./port_scanner -i 127.0.0.1 -p 80`
#### 	Oputput
```bash
Scanning port...
Results:
Port:80   State:Open   Protocol:tcp

```

# Arguments
##### You can do much more here is a list of command that you can use:
| Short form of argument (-)  | Long form of argument (- -)   |  Description|
| ------------ | ------------ | ------------ |
| i   |  ip |***Required*** specifies an ip that you like to scan. must be a valid ip
| p  | port  | specifies a port that you like to scan. must be a number between '0' and '65535'
| l  | protocol  |specifies the protocol. must be either 'tcp' or 'udp'  ***default value is 'tcp'***
| c  | common  | if 'true' scans common ports: '0' to '1024'  ***default value is 'false'***
| r  | range  | if 'true' asks for ranges ***default value is 'false'***
| d  | duration  | duration of connection. must be a value between '7' and '300'  ***default value is '10' ***

# Examples
#### Command
`$ ./port_scanner -i 127.0.0.1 -c true`
#### Output
```bash
Scanning ports...
Port:80   State:Open   Protocol:tcp
Port:443   State:Open   Protocol:tcp
Port:444   State:Open   Protocol:tcp
Port:631   State:Open   Protocol:tcp
Port:902   State:Open   Protocol:tcp

```
#### Command
`$ ./port_scanner -i 127.0.0.1 -r true`
#### Output
```bash
Enter the firts port number: 0
Enter the seconde port number: 1024
Scanning ports...
Port:80   State:Open   Protocol:tcp
Port:443   State:Open   Protocol:tcp
Port:444   State:Open   Protocol:tcp
Port:631   State:Open   Protocol:tcp
Port:902   State:Open   Protocol:tcp

```
####Command
`$ ./port_scanner -i 127.0.0.1 -c true -d 15 -l udp
`
##### runs the first example with a 15 seconde connection duration and using udp protocol.

###Notice
##### I'm fairly new to the golang world so the code is probably not so well written or efficient!!!
