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
