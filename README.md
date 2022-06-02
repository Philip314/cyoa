# Choose Your Own Adventure

## About

Web application of a story with a series of chapters where you can choose an option to go to the next chapter.

I created this from following [Gophercises](https://gophercises.com/ "Gophercises") to learn Go.

## Getting Started

1. Clone repo
2. Go into the project folder with command line interface
3. Run project with `go run cmd/cyoaweb/main.go`
4. Go to `localhost:8080/story/` in a web browser and enjoy the story

## Command-line Arguments

```
$ go run cmd/cyoaweb/main.go -help
 -file string
        JSON file of story (default "gopher.json")
 -port int
        Port that application runs on (default 8080)
```