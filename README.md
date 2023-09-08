# go-algo
This is a personal project. Probably not useful for others.

Learning go by implementing algorithms in it. 


## Make
I use make for developer tasks like building and installing.

Run `make` or `make help` to see what is available.


## Running the binary
The binary is `go-algo`. I alias it to `algo` in my bash aliases.

If you use `make build` then it will be available in project root folder, so you can run it like `./go-algo`. 

If you ran `make install` it will be available in path, so you can run it from anywhere on your machine like `$> go-algo`.



## Notes

- [2023-09-08] For watching filesystem and rebuilding and rerunning the commands, I tried out `watch make install` which works but is very basic. [`air`](https://github.com/cosmtrek/air) package has more features. So, I am using air. The command flow is: 
    - `make watch -> air [using config in .air.toml]  -> make install -> air's runner`

