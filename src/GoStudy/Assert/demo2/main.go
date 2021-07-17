package main

import "fmt"

type Phone struct {
}
type Camera struct {
}
type Computer struct {
}
type Usb interface {
	start()
	end()
}

func (phone Phone) start() {
	fmt.Println("phone start")
}
func (phone Phone) end() {
	fmt.Println("phone end")
}
func (phone Phone) call() {
	fmt.Println("phone call")
}

func (camera Camera) start() {
	fmt.Println("camera start")
}
func (camera Camera) end() {
	fmt.Println("camera end")
}

func (computer Computer) Working(usb Usb) {
	usb.start()
	if phone, ok := usb.(Phone); ok {
		phone.call()
	}
	usb.end()
}
func main() {

	usbArr := [3]Usb{}
	usbArr[0] = Phone{}
	usbArr[1] = Camera{}
	usbArr[2] = Phone{}
	c := Computer{}
	for _, val := range usbArr {
		c.Working(val)
		fmt.Println()
	}

}
