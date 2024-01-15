# goexpert_1
 
## Executing

### From the Binaries

There are some binaries built in the `binaries` folder.  
If you want to execute them you can use:

```shell
cd binaries
chmod +x server_<mac/linux>
chmod +x client_<mac/linux>

./server_<mac/linux>

# In another terminal
./client_<mac/linux>
```

### From the source code

You will need two terminals

#### Server
```shell
cd server
go run cmd/main.go
```

#### Client
```shell
cd client
go run cmd/main.go 
```

