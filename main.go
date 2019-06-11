/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	 if err := shim.Start(new(Chaincode)) ; err != nil{
	     fmt.Printf("Error starting chaincode %s" , err)
     }
}
