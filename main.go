package main

func main() {
	aTV := &tv{}
	anOnCommand := &onCommand{
		device: aTV,
	}
	anOffCommand := &offCommand{
		device: aTV,
	}
	anOnButton := &button{
		command: anOnCommand,
	}
	anOnButton.press()
	anOffButton := &button{
		command: anOffCommand,
	}
	anOffButton.press()
}
