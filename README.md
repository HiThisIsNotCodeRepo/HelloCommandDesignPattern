# HelloCommandDesignPattern

> [Source](https://golangbyexample.com/command-design-pattern-in-golang/)

## Core elements

```
type device interface {
	on()
	off()
}
```

`device` tells us how many functions a device should have, it also determines how many command object should we have, in
most simplified model we should have 2 command, one for `on()` and one for `off()`

```
type command interface {
	execute()
}
```

`command` have only `execute()` it is used to execute the device function.

```
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

```

`onCommand` and `offCommand` is to trigger device `on` and `off` function respectively.

```
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

```

`button` can be considered as first trigger of `command`, when client press the button it will trigger command execution
thus call the device function defined in `device`