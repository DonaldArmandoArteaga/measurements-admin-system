import { IVpc, SubnetType, Vpc } from "aws-cdk-lib/aws-ec2";
import { Construct } from "constructs";

export class InputSystemStackCommonInstances {
    private readonly defaultVPC: IVpc

    constructor(scope: Construct) {
        this.defaultVPC = new Vpc(scope, 'VPC', {
            natGateways: 0,
            subnetConfiguration: [
                {
                    cidrMask: 24,
                    name: "imput-system-default-vpc-public",
                    subnetType: SubnetType.PUBLIC
                },
                // {
                //     cidrMask: 24,
                //     name: "imput-system-default-vpc-private-isolated",
                //     subnetType: SubnetType.PRIVATE_ISOLATED
                // },
            ]
        });
    }

    get GeDefaultVPC(): IVpc {
        return this.defaultVPC
    }

}







