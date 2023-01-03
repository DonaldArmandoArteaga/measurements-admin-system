#!/bin/bash

sudo yum update
export QUEUE_NAME=raw-measurements-data-queue 
export DYNAMO_TABLE_NAME=raw-measurements-table
sudo yum install golang -y
cd home/ec2-user/
git clone https://github.com/DonaldArmandoArteaga/measurements-admin-system.git
cd measurements-admin-system/src/ec2/
sudo go build .
sudo go run input-system
