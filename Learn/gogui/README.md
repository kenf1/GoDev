## Go GUI App using Fyne

Steps:

```{shell}
#setup gogui folder
cd Learn
mkdir gogui
cd gogui
go mod init gogui
go mod tidy

#create empty main.go file
touch main.go

#install fyne
go install fyne.io/fyne/v2/cmd/fyne@latest

#run gui app
go run main.go
```
