# goProject

This template should help get you started.  
ListenAndServe 5000 (Example: localhost:5000)

## go installation
  1. Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:

```sh
$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz
```
You may need to run the command as root or through sudo).

Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.

2. Add /usr/local/go/bin to the PATH environment variable.
You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

```sh
export PATH=$PATH:/usr/local/go/bin
```
Note: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.

3. Verify that you've installed Go by opening a command prompt and typing the following command:

```sh
$ go version
```

### Compile Reload
```sh
go run main.go
```
example
```sh
$ goProject go run main.go
```
### API Router: chi  


### Resources
- [go](https://www.youtube.com/watch?v=lodF5BaZdDY&list=PLhdY0D_lA34W1wS2nJmQr-sssMDuQf-r8)
- [chi](https://github.com/go-chi/chi)
- [Dgraph go client](https://dgraph.io/docs/clients/go/)
- [dgo](https://pkg.go.dev/github.com/dgraph-io/dgo)
- [context](https://www.youtube.com/watch?v=yFN-IZQfdpY) 
- [Unmarshal](pkg.go.dev/encoding/json#example-Unmarshal)
