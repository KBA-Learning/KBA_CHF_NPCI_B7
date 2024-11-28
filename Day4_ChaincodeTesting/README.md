
# Start network using script file

```
./startNetwork.sh

```

# Minifab commands to deploy and invoke chaincode

```
sudo chmod -R 777 vars/
```
```
mkdir -p vars/chaincode/KBA-Automobile/go
```
```
cp -r ../Chaincode/* vars/chaincode/KBA-Automobile/go/
```
```
minifab ccup -n KBA-Automobile -l go -v 1.0 -d false -r false
```
```
minifab invoke -n KBA-Automobile -p '"CreateCar","car01","BMW","320d","Red","F-01","01/01/2024"'
```
```
minifab query -n KBA-Automobile -p '"ReadCar","car01"'

```

**Update the chaincode**
```
cp -r ../Chaincode/* vars/chaincode/KBA-Automobile/go/
```

```
minifab ccup -n KBA-Automobile -l go -v 1.1 -d false -r false
```

```
minifab invoke -n KBA-Automobile -p '"DeleteCar","car01"'
```
```
minifab query -n KBA-Automobile -p '"ReadCar","car01"'

```

**To stop network and restart it later**

```
minifab down
```
```
minifab restart
```

**To cleanup entire network**
```
minifab cleanup
```
```
sudo rm -rf vars
```

**Chaincode Tesing**

Create a new folder named test

```
mkdir test
```

Now create a file named car-contract_test.go inside test directory
```
touch test/car-contract_test.go
```

We need to create mocks for all the method for fabric communication
```
git clone https://github.com/maxbrunsfeld/counterfeiter.git
cd counterfeiter
go build
sudo cp counterfeiter /usr/local/go/bin

counterfeiter
```

Use the following command to generate mocks
```
cd test

go generate
```

Use this command to add the dependencies required for testing
```
go mod tidy
```

Use the following command to run the test files
```
go test
```

To see a detailed output run the following

```
go test -v

```
