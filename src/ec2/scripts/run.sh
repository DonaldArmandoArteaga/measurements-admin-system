#!/bin/bash

yum update
yum install golang -y
cd home/ec2-user/
git clone https://github.com/DonaldArmandoArteaga/measurements-admin-system.git
cd measurements-admin-system/src/ec2/
go build .
go run input-system