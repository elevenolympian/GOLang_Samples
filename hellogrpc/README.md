#### Compilation information

You have to compile protobuffer file (macOS)

```bash
protoc protobuff/messages.proto --go_out=plugins=grpc:go/src
```

In order to fix GO home problems 

```bash
  $  export GOPATH=$HOME/go
  $  export PATH=$GOPATH/bin:$PATH
  $  source ~/.bash_profile 
  ```

  Alternatively, you can do as below: 

  ```bash
$ export GOROOT=/usr/local/go
$ export GOPATH=$HOME/go
$ export GOBIN=$GOPATH/bin
$ export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
$ source ~/.bash_profile 
```


In order to eliminate GO111MODULE=auto, please type as below: 


```bash
$ go env
$ go env -w GO111MODULE=auto
$ go mod tidy
```

### Some important problems 

If you get an error as below: 

```bash
_cgo_export.c:3:10: fatal error: 'stdlib.h' file not found
```

Please type the following code snippet

```bash
$ export SDKROOT="$(xcrun --sdk macosx --show-sdk-path)"
```

And then you can build the go project with the following command

```bash
$ go build github.com/ps/hellogrpc
```

You can run the nodejs application with the following command

```bash
$ node nodejsapp/app.js
```

### Important note regarding the Nodejs application

If you run the application as below from the folder named **nodejsapp**: 

```bash
$ node app.js
```

You would get an error as below: 
```bash
Error: Could not load file "protobuff/messages.proto"
```

Because protobuff relative path is not same level with the nodejsapp folder so that you should run as below: 

```bash
node nodejsapp/app.js
```

Output is

```bash
Hello TU DRESDEN
````


