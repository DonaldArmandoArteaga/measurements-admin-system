import { IVpc, SubnetType, Vpc } from "aws-cdk-lib/aws-ec2";
import { Construct } from "constructs";

export class InputSystemStackCommonInstances {
    private readonly defaultVPC: IVpc

    constructor(scope: Construct) {
        // this.defaultVPC = new Vpc(scope, 'InputSystemVPC', {
        //     cidr: '10.0.0.0/16',
        //     natGateways: 0,
        //     maxAzs: 3,
        //     subnetConfiguration: [
        //         {
        //             name: 'public-subnet-1',
        //             subnetType: SubnetType.PUBLIC,
        //             cidrMask: 24,
        //         },
        //         {
        //             name: 'isolated-subnet-1',
        //             subnetType: SubnetType.PRIVATE_ISOLATED,
        //             cidrMask: 28,
        //         },
        //     ],
        // });

        this.defaultVPC = new Vpc(scope, 'VPC', {
            natGateways: 0,
            subnetConfiguration: [{
                cidrMask: 24,
                name: "asterisk",
                subnetType: SubnetType.PUBLIC
            }]
        });
    }

    get GeDefaultVPC(): IVpc {
        return this.defaultVPC
    }

}







