package src

import "github.com/jaypipes/ghw"

func GetGPU() string {
	// les get general info first
	gpu, err := ghw.GPU()
	if err != nil || len(gpu.GraphicsCards) == 0 {
		return "error tryin to fetch gpu name"
	}

	// lets get the first GPU card
	card := gpu.GraphicsCards[0]

	// query host's PCI database
	pci, err := ghw.PCI()
	if err != nil {
		return "Unknow GPU {PCI error}"
	}

	// Using cards hardware address we'll look at specific ddetails
	device := pci.GetDevice(card.Address)
	if err != nil {
		return "Unknow GPU {device not found}"
	}

	// lets get and combine Vendor and Product name
	vendorName := device.Vendor.Name
	productName := device.Product.Name

	return vendorName + " " + productName
}
