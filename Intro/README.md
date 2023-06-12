## Intro

```{shell}
#run file
go run intro.go

#build binary (current os + architecture)
go build intro.go

#run binary
./intro
```

```{shell}
#list all platforms to build for
go tool dist list

#build for windows amd64
GOOS=windows GOARCH=amd64 go build intro.go
```

### Platforms

|GOOS|GOARCH|Description|
|---|---|---|
|linux|amd64/arm64|
|windows|amd64/arm64|
|darwin|amd64/arm64|macOS|
