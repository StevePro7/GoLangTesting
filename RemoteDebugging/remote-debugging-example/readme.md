Remote Debugging
08-Feb-2023

https://kupczynski.info/posts/remote-debug-go-code
https://github.com/igor-kupczynski/remote-debug-example


ptrace
https://linux.die.net/man/2/ptrace


JetBrains
https://blog.jetbrains.com/go/2020/05/06/debugging-a-go-application-inside-a-docker-container/#Startingthedebugger
https://blog.jetbrains.com/go/tag/containers
https://blog.jetbrains.com/go/2020/05/11/using-kubernetes-from-goland


IMPORTANT
I had errors with initial go 1.14
go get dlv
https://github.com/apache/trafficcontrol/issues/6661
https://stackoverflow.com/questions/70816860/golang-build-specific-version-w-o-flag

Solution
go install
RUN go install github.com/go-delve/delve/cmd/dlv@2f13672765fe


Then I had issues with version of go and dlv compatibility
probably because of the git hash
2f13672765fe

so I just modified all go 1.19 to go 1.17 and docker worked OK


Finally, had to setup configuration as per image
Go Remote
port 40000

but when I did the curl to the debug version the code did not debug break
curl localhost:8081

So I clicked the debug button in Goland IDE and remote debug launched

Finally, the curl prompted the code to break into the source code
so I could debug step thru

curl localhost:8081