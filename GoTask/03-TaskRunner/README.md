Task Runner
01-Dec-2023

https://dev.to/ruanbekker/task-runner-with-yaml-config-written-in-go-kkm


Env var
environment variable


Global env var
version: '3'

env:
  WORD: world


Local env var
  byeworld:
    cmds:
      - echo "$GREETING, $WORD!"
    env:
      GREETING: bye


.env file
dotenv: ['.env']


Dependent tasks
  all:
    deps: [helloworld, byeworld

task all
Hello, there!
bye, world!
