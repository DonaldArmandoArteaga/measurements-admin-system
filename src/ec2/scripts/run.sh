#!/bin/bash


echo init-script-
sudo su
yum install git -y
cd home/ec2-user/
git clone https://github.com/DonaldArmandoArteaga/measurements-admin-system.git
cd measurements-admin-system/src/ec2/
yum install golang -y
export HOME=/home/ec2-user/
export GO111MODULE=on
sudo go build .
go run input-system
echo end-script-