//
// Copyright 2020 Yiwenlong(wlong.yi#gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package tps_starter

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
