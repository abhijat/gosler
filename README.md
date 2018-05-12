
### Gosler is a shell for Hashicorp vault

It is a port to golang of the rust project [mosler](https://github.com/abhijat/mosler)

The goal is to be able to login once when starting, and then run commands which perform operations using the vault HTTP api. 

The auth token will be cached for the login session and will be used to communicate with the vault server.

<br>

###### Dependencies

* github.com/peterh/liner


