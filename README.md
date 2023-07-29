### Internet Cafe System
a small app for automate booking pc in internet cafe with interactive command in CLI

Create the binary file from:
```
go build ./bin/main.go
```

Supported Command:
```
create_internet_cafe <total pc capacity number>
status
book <user name> <user age>
leave <pc number>
get_pc_by_age <user age>
```