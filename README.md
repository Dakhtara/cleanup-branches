
# GIT Cleanup branches

The purpose of this repo is to remove all stale branches and all no remote branch local branches. 

This code is mainly for training with Go. 


## Build pour MacOS
```sh
go build cleanup-branches
```


## Build pour linux

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cleanup-branches-lin
```

## Contributing

If you found a bug or want to improve the codebase feel free to open a Pull Request.
