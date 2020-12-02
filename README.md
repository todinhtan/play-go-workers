# play-go-workers
Here is a quick sample of using goroutine, channel and waitgroup together
The application will summon some workers to make a list of string uppercase
1. Start the tcp server at port 8080, with a maximum workers to do the tasks
```
go run main.go X
```
2. Connect to tcp server (using telnet)
```
telnet localhost 8080
```
3. Provide a list of string seperated by a comma
```
thor,hulk,black widow,iron man,ant man
```
4. Cheers