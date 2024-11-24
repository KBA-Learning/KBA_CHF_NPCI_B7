# Installing the Dependencies

Note: If any of the following dependencies are available on your laptop, then no need to install it.

## Update Packages

In case of a fresh Ubuntu 22 installation, use the following commands to update the packages before installing other dependencies.  
```
sudo apt update
```

```
sudo apt upgrade
```

## Visual Studio Code
Download and install the latest version of VS code from here: https://code.visualstudio.com/download


To install, execute the following command from the same folder where VS Code is being downloaded.

Note: Replace file_name with the actual name of the file you've downloaded.
```
sudo dpkg -i file_name
```
eg: sudo dpkg -i code_1.95.2-1730981514_amd64.deb


## cURL
Install curl using the command
```
sudo apt install curl -y
```

```
curl -V
```

## NVM

Install NVM (Node Version Manager), open a terminal and execute the following command.
```
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash
```
Close the current terminal and open a new one.

In the new terminal execute this command to verify nvm has been installed

```
nvm -v
```


## NodeJS (Ver 22.x)

Execute the following command to install NodeJs
```
nvm install 22
```  

Check  the version of nodeJS installed
```
node -v
```

Check  the version of npm installed
```
npm -v
```

## Docker
Step 1: Download the script
```
curl -fsSL https://get.docker.com -o install-docker.sh
```
Step 2: Run the script either as root, or using sudo to perform the installation.
```
sudo sh install-docker.sh
```
Step 2: To manage Docker as a non-root user
```
sudo chmod 777 /var/run/docker.sock
```

```
sudo usermod -aG docker $USER
```

To verify the installtion enter the following commands


```
docker compose version
```

```
docker -v
```

## JQ
Install JQ using the following command
```
sudo apt install jq -y
```

To verify the installtion enter the following command


```
jq --version
```

## Build Essential
Install Build Essential uisng the commnad
```
sudo apt install build-essential -y
```
To verify the installtion enter the following command


```
dpkg -l | grep build-essential
```

## Go
Step 1: Download Go
```
curl -OL  https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
```
Step 2: Extract
```
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
```

Step 3: Add /usr/local/go/bin to the PATH environment variable. Open the /etc/environment file
```
sudo gedit /etc/environment
```

Step 4: Append the following to the end of `PATH` variable and save
```
:/usr/local/go/bin
```
```
source $HOME/.profile
```

To verify the installtion enter the following command


```
go version
```
Note: If go version is not listed, then restart the system and execute the command again.

## Minifab

```
curl -o minifab -sL https://tinyurl.com/yxa2q6yr && chmod +x minifab
```

```
sudo cp minifab /usr/local/bin
```

```
minifab
```

## Git
Install git using the command
```
sudo apt install git -y
```
To verify the installtion enter the following command
```
git --version
```

**Create GitHub account**: Browse https://github.com/, sign up by entering email id, password and user name. Then solve the puzzle and enter the GitHub launch code recieved on your email. Select the account preferences and create a free account. 

**Set Up Git**: Configure your Git username and email using the following commands:
```bash
git config --global user.name "Your Github UserName"
```
```
git config --global user.email "your Github email"
```

**SSH Key Configuration**
```
ssh-keygen -t ed25519 -C "github_email" //Replace the emaial with any recognizable tag that you like
```
```
eval "$(ssh-agent -s)"
```
```
ssh-add ~/.ssh/id_ed25519
```
```
cat ~/.ssh/id_ed25519.pub
```

Copy the key from ssh to the last.

Now, in github account, go to the settings and select SSH and GPG keys, click New SSH key, enter a title, paste this key and click add.

Now if you clone using ssh. It will automatically be cloned to your local sytem.


#### Git Commands

1. Initialize git in a directory/folder
```bash
git init
```
2. Stage the changes

use this to add by file name
```bash
git add file_name
```
use this for adding all the changes to staging area
```
git add . 
```
3. Commit the changes
```bash
git commit -m "commit_message"
```

4. Push to the repo
```bash
git push
```
5. Pull from the remote repo
```bash
git pull 
```
6. Clone a repo (HTTPS or SSH)
```bash
git clone repository_url 
```