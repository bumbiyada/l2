package main

import "fmt"

// sender
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

// interface

type command interface {
	execute()
}

//command

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

//command

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

// interface
type device interface {
	on()
	off()
}

// getter

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

//client

func main() {
	tv := &tv{}

	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	onButton := &button{
		command: onCommand,
	}
	onButton.press()

	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
