# Go command

- build:

  ```shell
  go build
  ``` 
  

# How to test

## simulate TCP client connection

```shell
nc 127.0.0.1 8888
```


# API
## how to send msg?


0.

   ```go
   go func(){
     for{
        msg := <-user.chan
     }
   }()
   ```
   
1. `server.chan <- msg`
2.
   ```go
   for _, u:= range server.onlineuser{
        u.chan <- msg
   }
   ```