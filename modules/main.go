package main

import (
	"github.com/donvito/hellomod"
	hellomodv2 "github.com/donvito/hellomod/v2"
)

// Descargar de módulos de terceros, alias y varias versiones del mismo módulo

func main() {
	hellomod.SayHello()
	hellomodv2.SayHello("SantiSite ")
}
