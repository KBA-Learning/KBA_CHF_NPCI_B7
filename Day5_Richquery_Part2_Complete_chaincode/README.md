#### Minifab commands to deploy and invoke chaincode 

Note: Execute the following commands from Minifab_Network folder.


`./startNetwork.sh`

`sudo chmod -R 777 vars/`

`mkdir -p vars/chaincode/KBA-Automobile/go`

`cp -r ../Chaincode/* vars/chaincode/KBA-Automobile/go/`

`cp vars/chaincode/KBA-Automobile/go/collection-minifab.json ./vars/KBA-Automobile_collection_config.json`


`minifab ccup -n KBA-Automobile -l go -v 1.0 -d false -r true`



`minifab invoke -n KBA-Automobile -p '"CreateCar","car01","Tata","Tiago","White","F-01","22/03/2023"'`

`minifab invoke -n KBA-Automobile -p '"CreateCar","car02","Tata","Punch","Red","F-01","22/04/2022"'`

`minifab invoke -n KBA-Automobile -p '"CreateCar","car03","Honda","Amaze","Blue","F-03","18/03/2023"'`

`minifab invoke -n KBA-Automobile -p '"CreateCar","car04","Tata","Tiago","Red","F-01","22/03/2023"'`

`minifab invoke -n KBA-Automobile -p '"CreateCar","car05","Tata","Altroz","Black","F-05","10/02/2023"'`



`minifab query -n KBA-Automobile -p '"GetCarHistory","car01"'`

`minifab query -n KBA-Automobile -p '"GetAllCars"'`

`minifab query -n KBA-Automobile -p '"GetCarsWithPagination","3",""'`

`minifab query -n KBA-Automobile -p '"GetCarsWithPagination","2","g1AAAAA6eJzLYWBgYMpgSmHgKy5JLCrJTq2MT8lPzkzJBYqzJicWGRiDJDlgkgjhLADJBA_x"'`


```
MAKE=$(echo -n "Tata" | base64 | tr -d \\n)
MODEL=$(echo -n "Tiago" | base64 | tr -d \\n)
COLOR=$(echo -n "White" | base64 | tr -d \\n)
DEALER_NAME=$(echo -n "XXX" | base64 | tr -d \\n)

```
`minifab invoke -n KBA-Automobile -p '"OrderContract:CreateOrder","ord01"' -t '{"make":"'$MAKE'","model":"'$MODEL'","color":"'$COLOR'","dealerName":"'$DEALER_NAME'"}' -o dealer.auto.com`

```
MAKE=$(echo -n "Tata" | base64 | tr -d \\n)
MODEL=$(echo -n "Tiago" | base64 | tr -d \\n)
COLOR=$(echo -n "White" | base64 | tr -d \\n)
DEALER_NAME=$(echo -n "XXX" | base64 | tr -d \\n)

```

`minifab invoke -n KBA-Automobile -p '"OrderContract:CreateOrder","ord03"' -t '{"make":"'$MAKE'","model":"'$MODEL'","color":"'$COLOR'","dealerName":"'$DEALER_NAME'"}' -o dealer.auto.com`

`minifab query -n KBA-Automobile -p '"OrderContract:ReadOrder","ord01"'`



#### After updating chaincode Execute following commands from Minifab_Network folder

`cp -r ../Chaincode/* vars/chaincode/KBA-Automobile/go/`

`minifab ccup -n KBA-Automobile -l go -v 2.0 -d false -r true`

`minifab query -n KBA-Automobile -p '"GetMatchingOrders","car01"'`

`minifab invoke -n KBA-Automobile -p '"MatchOrder","car01","ord01"'`

`minifab invoke -n KBA-Automobile -p '"RegisterCar","car01","Bob","KL-01-XXXX"' -o mvd.auto.com`

`minifab query -n KBA-Automobile -p '"ReadCar","car01"'`

`minifab query -n KBA-Automobile -p '"GetCarHistory","car01"'`



