# UI

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

## Create the `Simple-App` folder

```bash
cd ..
```

```bash
mkdir Simple-App
```

```bash
cd Simple-App
```
### Build the application


```bash
go mod init simple-app
```

Copy profile.go, connect.go and client.go to this folder

```bash
cp ../Client/profile.go ../Client/connect.go ../Client/client.go ./
```
```bash
go get google.golang.org/grpc@v1.67.1
```
Create the `main.go` in `Simple-App` folder, `index.html` in `templates` folder and index.js in `public/scripts` folder

```bash
go mod tidy
```
```bash
go run .
```

Browse http://localhost:3000/

To stop, execute `ctrl+c`


To run the complete application, navigate to `Auto-App` folder,

 
```bash
cd Auto-App/
```
```bash
go mod tidy
```
```bash
go run .
```

Browse http://localhost:8080/

To stop, execute `ctrl+c`

### Stop the automobile network

```bash
cd Automobile-network/
```

To start the automobile network execute this command

```bash
./stopAutomobileNetwork.sh
```
