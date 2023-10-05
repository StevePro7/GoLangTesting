Remove the go then there are no goroutines running
ignore the main goroutine


Step interrupted by a breakpoint. Use 'Continue' to resume the original step command.
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.boring({0x4b67a8, 0x6}, 0xc000126060)
	/home/stevepro/GitHub/StevePro7/GoLangTesting/ConcurrencyTest/RobPike/03_Chan/main.go:11 +0x145
main.main()
	/home/stevepro/GitHub/StevePro7/GoLangTesting/ConcurrencyTest/RobPike/03_Chan/main.go:18 +0x4d