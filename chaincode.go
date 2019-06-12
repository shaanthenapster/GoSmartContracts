package main

import (
    "encoding/json"
    "fmt"
    _ "github.com/hyperledger/fabric/bccsp/utils"
    "math/rand"
    "time"
    _ "math/rand"
    "github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}

type User struct {

    UserID string `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
    MobileNo string `json:"mobile_no"`
}

type Land struct {
    LandId string `json:"land_id"`
    LandDescription string `json:"land_description"`
    Owner User `json:"owner"`
    CreatedAt time.Time `json:"created_at"`
}

func (s *Chaincode) createUser(stub shim.ChaincodeStubInterface) sc.Response{

    args := stub.GetStringArgs()
    Id := string(rand.Int())
    var user  = User{UserID:Id , Name:args[0], Email: args[1] , MobileNo: args[3] }
    UserBytes , _ := json.Marshal(user)
    stub.PutState(Id , UserBytes)
    return shim.Success(nil)
}

func (s *Chaincode) uploadLand(stub shim.ChaincodeStubInterface , user User) sc.Response  {
     args := stub.GetStringArgs()
     LandId := "RAW-LAND" + string(rand.Int())
     var land = Land{LandId:LandId , LandDescription: args[0] , Owner:user}
     LandBytes , _ := json.Marshal(land)
     stub.PutState(LandId , LandBytes)
     return shim.Success(nil)
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	args := stub.GetStringArgs()
	fmt.Println("I am argument 1", args[0])
	err := stub.PutState(args[1], []byte(args[2]))
	if err != nil {
		shim.Error(fmt.Sprintf("Failed to Create Asset: %s", args[1]))
	}
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fn, args := stub.GetFunctionAndParameters()
	var result string
	var err error
	if fn == "set" {
		result, err = set(stub, args)
	} else {
		result, err = get(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(result))
}

func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {

	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect Arguments")
	}
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return args[1], nil
}

func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return string(value), nil
}
