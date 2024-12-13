### Fabric Test Network commands 

**Note: Open a terminal in the KBA-CHF Folder & Execute the Following Commands**

``` bash
curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh
```
Download the script to install the fabric binaries, fabric images and fabric samples

``` bash
./install-fabric.sh -f '2.5.4' -c '1.5.7'
```
Install the version 2.5.4 and ca version 1.5.7

``` bash
sudo cp fabric-samples/bin/* /usr/local/bin
```
Copy the fabric binaries to usr/local/bin to use it in the whole system.

## To use the script navigate to test-network folder inside the fabric-samples folder,

`cd fabric-samples/test-network/`

`./network.sh -h`

`./network.sh up createChannel`

`docker ps -a`

### Deploy asset-transfer basic sample chaincode listed in the samples.

Note: Edit the go.mod file at `fabric-samples/asset-transfer-basic/chaincode-go`, change `go 1.22.0` to `go 1.22`

`./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go -ccl go`

```
export FABRIC_CFG_PATH=$PWD/../config/


export CORE_PEER_TLS_ENABLED=true


export CORE_PEER_LOCALMSPID="Org1MSP"

export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp

export CORE_PEER_ADDRESS=localhost:7051

```

`peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'`


`peer chaincode query -C mychannel -n basic -c '{"Args":["GetAllAssets"]}'`


`peer chaincode query -C mychannel -n basic -c '{"function":"ReadAsset","Args":["asset5"]}'`

### Stopping the network

`./network.sh down`