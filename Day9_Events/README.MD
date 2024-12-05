**#Create a Events folder with in the KBA-Automobile directory**

```

mkdir Events

```

**#Build the events code**

```
cd Events

go mod init events

Create & Build profile.go, connect.go, events.go, main.go

go mod tidy



```

### Set a car creation event in the CreateCar function in car-contract.go 

```
type EventData struct{
	Type string
	Model string
}
```


```
			addCarEventData := EventData{
				Type: "Car creation",
				Model: model,
			}
			eventDataByte, _ := json.Marshal(addCarEventData)
			ctx.GetStub().SetEvent("CreateCar",eventDataByte)
```



## Start the automobile network**

```
cd ..

cd Automobile-network/

```



To start the automobile network execute this command


```
./startAutomobileNetwork.sh

```


### To Run Block Event



Open a terminal in the events folder & consider this terminal as to listening block events.

`go run .`

Note: For checking newly created block do a transaction using client application.(open a terminal from Client folder and execute `go run .`)

### To Run Chaincode Event

open another terminal from Events folder and consider it as to listening chaincode events.

`go run .`


Note: Do a create car transaction using client application.( change CreateCar transaction arguments in main.go then execute `go run .` in the client terminal)

### To Run Private Block Event

Open a new terminal from Events folder and consider it to listening private blockevent.

`go run .`

Note: Submit a CreateOrder transaction using client application.(Edit main.go for CreateOrder transaction and execute `go run .`in the client terminal)


**#To stop the Automobile network**

```
./stopAutomobileNetwork.sh

```
