## Basic CLI

Create basic/simple cli app using Cobra.

### Steps:

***Note*** Run in bash

```{shell}
#run in bash
bash

#install packages (cobra + viper)
go get -u github.com/spf13/cobra@latest
# go get -u github.com/spf13/viper@v1.8.1

#create folder & use as working dir
mkdir basic_cli
cd basic_cli

#create go.mod file
go mod init

#setup cobra cli app
cobra-cli init
```

### Run basic_cli app

```{shell}
#run root cmd
go run main.go

#run get cmd
go run main.go get
```
