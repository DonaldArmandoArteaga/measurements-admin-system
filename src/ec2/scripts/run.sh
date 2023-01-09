#!/bin/bash

sudo yum update
sudo yum install golang -y
cd home/ec2-user/
git clone https://github.com/DonaldArmandoArteaga/measurements-admin-system.git
cd measurements-admin-system/src/ec2/
sudo go build .
sudo go run input-system
