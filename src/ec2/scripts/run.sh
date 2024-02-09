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
export GIN_MODE=release
echo $QUEUE_NAME
echo -------------------
export AA=$API_URL 
echo $AA
sudo go build .
echo ___________________
nohup 'export QUEUE_NAME=$QUEUE_NAME && export DYNAMO_TABLE_NAME=$DYNAMO_TABLE_NAME && export AWS_REGION=$AWS_REGION sudo ./input-system &'
echo input-system-background-process
cd /home/ec2-user/
git clone https://github.com/DonaldArmandoArteaga/measurements-mock
cd measurements-mock
sudo go build .
nohup 'export API_URL=$API_URL && sudo ./mesasurements-mock &'
echo end-script