# Client

**KBA-CHF -> KBA-Automobile -> Automobile-network**

**Open a command terminal with in Automobile-network folder**

```bash
cd Automobile-network/
```

### Start the automobile network

To start the automobile network execute this command

```bash
./startAutomobileNetwork.sh
```

## Create the client folder

```bash
cd ..
```

```bash
mkdir Client
```

### Build the client application

```bash
cd Client

go mod init client

Create and build profile.go, connect.go, client.go, main.go

go get google.golang.org/grpc@v1.67.1

go mod tidy

go run .

```

### To register and enroll a new user

**Execute these commands in the Automobile-network terminal**

```bash
export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/manufacturer.auto.com/

fabric-ca-client register --caname ca-manufacturer --id.name user2 --id.secret user2pw --id.type client --tls.certfiles "${PWD}/organizations/fabric-ca/manufacturer/ca-cert.pem"

fabric-ca-client enroll -u https://user2:user2pw@localhost:7054 --caname ca-manufacturer -M "${PWD}/organizations/peerOrganizations/manufacturer.auto.com/users/User2@manufacturer.auto.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/manufacturer/ca-cert.pem"  

```

## Stop the automobile network

```bash
./stopAutomobileNetwork.sh
```