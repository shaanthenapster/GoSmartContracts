/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
    sc "github.com/hyperledger/fabric/protos/peer"
)

// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {

    args := stub.GetStringArgs()
    if len(args) != 2{
        return shim.Error("Incorrect Arguments , Expecting a key Pair Value")
    }

    err := stub.PutState(args[0] , []byte(args[1]))
    if err!=nil{
        shim.Error(fmt.Sprintf("Failed to Create Asset: %s" , args[0]))
    }
    return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fn, args := stub.GetFunctionAndParameters()
	var result string
	var err error
	if fn == "set"{
	    result , err = set(stub , args)
    }else {
        result ,err = get(stub , args)
    }
	if err != nil{
	    return shim.Error(err.Error())
    }
	return shim.Success([] byte(result))

	fmt.Println("Invoke()", fcn, params)
	return shim.Success(nil)
}

func set()  {
    
}

