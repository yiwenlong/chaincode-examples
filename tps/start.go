package tps

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func BootChaincode(ccid, address string) {
	server := &shim.ChaincodeServer{
		CCID:    ccid,
		Address: address,
		CC:      new(Tps),
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}
	if err := server.Start(); err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
