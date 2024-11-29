**#Bring up the Minifab Network**

Go to the folder where you have Minifab script file and execute this command 

```
chmod +x startNetwork.sh

./startNetwork.sh
```

**#Execute these commands to deploy the chaincode with PDC:**

```
sudo chmod -R 777 vars/

mkdir -p vars/chaincode/KBA-Automobile/go

cp -r ../Chaincode/* vars/chaincode/KBA-Automobile/go/

cp vars/chaincode/KBA-Automobile/go/collection-minifab.json ./vars/KBA-Automobile_collection_config.json

minifab ccup -n KBA-Automobile -l go -v 1.0 -d false -r true

minifab invoke -n KBA-Automobile -p '"CreateCar","car01","Tata","Tiago","White","F-01","22/07/2024"'

minifab query -n KBA-Automobile -p '"ReadCar","car01"'

MAKE=$(echo -n "Tata" | base64 | tr -d \\n)

MODEL=$(echo -n "Tiago" | base64 | tr -d \\n)

COLOR=$(echo -n "White" | base64 | tr -d \\n)

DEALER_NAME=$(echo -n "XXX" | base64 | tr -d \\n)

minifab invoke -n KBA-Automobile -p '"OrderContract:CreateOrder","ord01"' -t '{"make":"'$MAKE'","model":"'$MODEL'","color":"'$COLOR'","dealerName":"'$DEALER_NAME'"}' -o dealer.auto.com

minifab query -n KBA-Automobile -p '"OrderContract:ReadOrder","ord01"'
```

**Bring down the network**

`minifab cleanup`



