### Update and upgrade all packages
echo "--- Upgrade and Update ---"
sudo apt-get update
sudo apt-get -y upgrade

### Install Go
echo "--- Installing Go ---"
# Change to Apps directory for all app resource download
echo "Downloading, extracting and linking Go"
mkdir -p /vagrant/resources/apps
cd /vagrant/resources/apps
sudo wget --no-verbose --timestamping https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz
mkdir -p go1.14.2
sudo tar --extract --skip-old-files --file /vagrant/resources/apps/go1.14.2.linux-amd64.tar.gz -C ./go1.14.2

# mkdir -p /usr/local/go is not required as detailed here:
# https://askubuntu.com/questions/600714/creating-a-symlink-from-one-folder-to-another-with-different-names/600732
sudo ln --symbolic --force /vagrant/resources/apps/go1.14.2/go /usr/local/go

## Go Settings


sudo echo ' ' >> /etc/profile.d/go-setup.sh
# 1. Set up bin location for Go
# as mentioned here: https://golang.org/doc/install
echo "Setting up bin location for Go"
sed -e 's|export PATH=$PATH:/usr/local/go/bin||g' -i /etc/profile.d/go-setup.sh
sudo echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile.d/go-setup.sh
#echo $PATH

# 2. Set up Workspace directory (GOPATH)
# as mentioned here: https://golang.org/doc/code.html#Workspaces
echo "Setting up GOPATH"
sed -e 's|export GOPATH=/vagrant/go||g' -i /etc/profile.d/go-setup.sh
sudo echo 'export GOPATH=/vagrant/go' >> /etc/profile.d/go-setup.sh
#echo $GOPATH

# 3. Set up Workspace bin dir
echo "Setting up bin location for GOPATH"
sed -e 's|export PATH=$PATH:$GOPATH/bin||g' -i /etc/profile.d/go-setup.sh
sudo echo 'export PATH=$PATH:$GOPATH/bin' >> /etc/profile.d/go-setup.sh
#echo $PATH