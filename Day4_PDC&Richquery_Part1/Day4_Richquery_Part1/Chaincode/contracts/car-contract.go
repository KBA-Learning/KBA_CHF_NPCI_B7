package contracts

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CarContract contract for managing CRUD for Car
type CarContract struct {
	contractapi.Contract
}

type Car struct {
	AssetType         string `json:"assetType"`
	CarId             string `json:"carId"`
	Color             string `json:"color"`
	DateOfManufacture string `json:"dateOfManufacture"`
	Make              string `json:"make"`
	Model             string `json:"model"`
	OwnedBy           string `json:"ownedBy"`
	Status            string `json:"status"`
}

// CarExists returns true when asset with given ID exists in world state
func (c *CarContract) CarExists(ctx contractapi.TransactionContextInterface, carID string) (bool, error) {
	data, err := ctx.GetStub().GetState(carID)

	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return data != nil, nil
}

// CreateCar creates a new instance of Car
func (c *CarContract) CreateCar(ctx contractapi.TransactionContextInterface, carID string, make string, model string, color string, manufacturerName string, dateOfManufacture string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()

	if err != nil {
		return "", fmt.Errorf("could not fetch client identity. %s", err)
	}

	if clientOrgID == "manufacturer-auto-com" {

		exists, err := c.CarExists(ctx, carID)
		if err != nil {
			return "", fmt.Errorf("could not fetch the details from world state.%s", err)
		} else if exists {
			return "", fmt.Errorf("the car, %s already exists", carID)
		}

		car := Car{
			AssetType:         "car",
			CarId:             carID,
			Color:             color,
			DateOfManufacture: dateOfManufacture,
			Make:              make,
			Model:             model,
			OwnedBy:           manufacturerName,
			Status:            "In Factory",
		}

		bytes, _ := json.Marshal(car)

		err = ctx.GetStub().PutState(carID, bytes)
		if err != nil {
			return "", fmt.Errorf("could not create car. %s", err)
		} else {
			return fmt.Sprintf("successfully added car %v", carID), nil
		}

	} else {
		return "", fmt.Errorf("user under following MSPID: %v can't perform this action", clientOrgID)
	}
}

// ReadCar retrieves an instance of Car from the world state
func (c *CarContract) ReadCar(ctx contractapi.TransactionContextInterface, carID string) (*Car, error) {

	bytes, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if bytes == nil {
		return nil, fmt.Errorf("the car %s does not exist", carID)
	}

	var car Car

	err = json.Unmarshal(bytes, &car)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type Car")
	}

	return &car, nil
}

// DeleteCar removes the instance of Car from the world state
func (c *CarContract) DeleteCar(ctx contractapi.TransactionContextInterface, carID string) (string, error) {

	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", fmt.Errorf("could not fetch client identity. %s", err)
	}
	if clientOrgID == "manufacturer-auto-com" {

		exists, err := c.CarExists(ctx, carID)
		if err != nil {
			return "", fmt.Errorf("%s", err)
		} else if !exists {
			return "", fmt.Errorf("the car, %s does not exist", carID)
		}

		err = ctx.GetStub().DelState(carID)
		if err != nil {
			return "", err
		} else {
			return fmt.Sprintf("car with id %v is deleted from the world state.", carID), nil
		}

	} else {
		return "", fmt.Errorf("user under following MSPID: %v can't perform this action", clientOrgID)
	}
}

func (c *CarContract) GetCarsByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*Car, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the private data by range. %s", err)
	}
	defer resultsIterator.Close()

	return carResultIteratorFunction(resultsIterator)
}

func (c *CarContract) GetAllCars(ctx contractapi.TransactionContextInterface) ([]*Car, error) {

	queryString := `{"selector":{"assetType":"car"}, "sort":[{ "color": "desc"}]}`

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("could not fetch the query result. %s", err)
	}
	defer resultsIterator.Close()
	return carResultIteratorFunction(resultsIterator)
}

func carResultIteratorFunction(resultsIterator shim.StateQueryIteratorInterface) ([]*Car, error) {
	var cars []*Car
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("could not fetch the details of the result iterator. %s", err)
		}
		var car Car
		err = json.Unmarshal(queryResult.Value, &car)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal the data. %s", err)
		}
		cars = append(cars, &car)
	}

	return cars, nil
}
