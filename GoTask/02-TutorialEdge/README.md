Tutorial Edge
01-Dec-2023

https://tutorialedge.net/golang/taskfiles-for-go-developers

task --list-all
task: Available tasks for this project:
* hello:

Now with desc
task --list
task: Available tasks for this project:
* hello:       A Hello World Task


Subdir
dir: terraform
task terraform-apply
/home/stevepro/GitHub/StevePro7/GoLangTesting/GoTask/02-TutorialEdge/terraform


Task dependencies
deps: [terraform-plan]
task terraform-apply
Running terraform plan...
Running terraform apply...


Dynamic Variables
vars
e.g.
{{.HOMEDIR}}
task install
Installing tool into /home/stevepro/.mydir


General Taskfile for Go devs
build       go build -o app cmd/server/main.go
test        go test -v ./...
lint        golangci-lint run
run         docker-compose up --build
int         go test -tags=integration -v ./...
