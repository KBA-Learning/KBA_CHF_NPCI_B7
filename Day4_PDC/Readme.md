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

**Execute the following command to install the additional shim package:**

Execute this command within the chaincode folder

`go mod tidy`


**#Deploy the chaincode again to use rich queries**

Execute this command in the Minifab folder:

```
cp -r ../Chaincode/* vars/chaincode/KBA-Automobile/go/

minifab ccup -n KBA-Automobile -l go -v 2.0 -d false -r true
```


**#Create additional cars to execute the rich queries**

```
minifab invoke -n KBA-Automobile -p '"CreateCar","car02","Maruti","Swift","Red","F-01","25/06/2024"' -o manufacturer.auto.com

minifab invoke -n KBA-Automobile -p '"CreateCar","car03","Kia","Seltos","Black","F-01","10/08/2024"'

minifab invoke -n KBA-Automobile -p '"CreateCar","car04","Honda","Amaze","Yellow","F-01","01/11/2024"'

minifab invoke -n KBA-Automobile -p '"CreateCar","car05","Hyundai","Creta","Blue","F-01","18/09/2024"'

minifab query -n KBA-Automobile -p '"GetAllCars"'

minifab query -n KBA-Automobile -p '"GetCarsByRange","car01","car05"'
```

**#Execute the following commands to sort the Cars**

Create the folder structure within Chaincode folder

```
mkdir -p META-INF/statedb/couchdb/indexes
```

ADD the inedxColor.json file to `indexes` folder


```
cp -r ../Chaincode/* vars/chaincode/KBA-Automobile/go/

minifab ccup -n KBA-Automobile -l go -v 3.0 -d false -r true

minifab query -n KBA-Automobile -p '"GetAllCars"'
```

**#Create additional orders to execute the rich queries**

```
MAKE=$(echo -n "Maruti" | base64 | tr -d \\n)

MODEL=$(echo -n "Swift" | base64 | tr -d \\n)

COLOR=$(echo -n "Red" | base64 | tr -d \\n)

DEALER_NAME=$(echo -n "XXX" | base64 | tr -d \\n)

minifab invoke -n KBA-Automobile -p '"OrderContract:CreateOrder","ord02"' -t '{"make":"'$MAKE'","model":"'$MODEL'","color":"'$COLOR'","dealerName":"'$DEALER_NAME'"}' -o dealer.auto.com

MAKE=$(echo -n "Kia" | base64 | tr -d \\n)

MODEL=$(echo -n "Seltos" | base64 | tr -d \\n)

COLOR=$(echo -n "Black" | base64 | tr -d \\n)

DEALER_NAME=$(echo -n "XXX" | base64 | tr -d \\n)

minifab invoke -n KBA-Automobile -p '"OrderContract:CreateOrder","ord03"' -t '{"make":"'$MAKE'","model":"'$MODEL'","color":"'$COLOR'","dealerName":"'$DEALER_NAME'"}' -o dealer.auto.com

MAKE=$(echo -n "Honda" | base64 | tr -d \\n)

MODEL=$(echo -n "Amaze" | base64 | tr -d \\n)

COLOR=$(echo -n "Yellow" | base64 | tr -d \\n)

DEALER_NAME=$(echo -n "XXX" | base64 | tr -d \\n)

minifab invoke -n KBA-Automobile -p '"OrderContract:CreateOrder","ord04"' -t '{"make":"'$MAKE'","model":"'$MODEL'","color":"'$COLOR'","dealerName":"'$DEALER_NAME'"}' -o dealer.auto.com

MAKE=$(echo -n "Hyundai" | base64 | tr -d \\n)

MODEL=$(echo -n "Creta" | base64 | tr -d \\n)

COLOR=$(echo -n "Blue" | base64 | tr -d \\n)

DEALER_NAME=$(echo -n "XXX" | base64 | tr -d \\n)

minifab invoke -n KBA-Automobile -p '"OrderContract:CreateOrder","ord05"' -t '{"make":"'$MAKE'","model":"'$MODEL'","color":"'$COLOR'","dealerName":"'$DEALER_NAME'"}' -o dealer.auto.com

minifab query -n KBA-Automobile -p '"OrderContract:GetAllOrders"'

minifab query -n KBA-Automobile -p '"OrderContract:GetOrdersByRange","ord01","ord05"'

```

