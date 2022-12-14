package fedex

import "fmt"

type FSender struct {
	Sender string
}

func (f *FSender) Send(parcel string) {
	fmt.Printf("Fedex sends %v parcel\n", parcel)
}
