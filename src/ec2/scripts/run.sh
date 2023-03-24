#!/bin/bash


echo init-script
sudo su
yum install git -y
cd home/ec2-user/
git clone https://github.com/DonaldArmandoArteaga/measurements-admin-system.git
cd measurements-admin-system/src/ec2/
yum install golang -y
export HOME=/home/ec2-user/
export GO111MODULE=on
sudo go build .
nohup sudo ./input-system &
echo input-system-background-process
cd /home/ec2-user/
git clone https://github.com/DonaldArmandoArteaga/measurements-mock
cd measurements-mock
sudo go build .
nohup sudo ./mesasurements-mock & 
echo end-script-