https://github.com/hyperledger-labs/minifabric

https://github.com/hyperledger-labs/minifabric/blob/main/spec.yaml


####    Bring up the network ###

Note: Execute the following commands from Minifab_Network folder, where the spec.yaml file is available.

```
minifab netup -s couchdb -e true -i 2.4.8 -o manufacturer.auto.com
```
```
minifab create -c autochannel
```
```
minifab join -c autochannel
```
```
minifab anchorupdate
```
```
minifab profilegen
```
### Bring down the network ###
```
minifab cleanup
```
### Using script ###
```
chmod +x startNetwork.sh
```
```
./startNetwork.sh
```

