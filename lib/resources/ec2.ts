import { CfnOutput } from "aws-cdk-lib"
import { Vpc, SecurityGroup, Peer, Port, UserData, Instance, InstanceType, InstanceClass, InstanceSize, MachineImage, AmazonLinuxGeneration, CloudFormationInit, InitFile, InitConfig, IVpc } from "aws-cdk-lib/aws-ec2"
import { Role, ServicePrincipal } from "aws-cdk-lib/aws-iam"
import { Construct } from "constructs"


export class InputSystemStackEC2 {
    constructor(scope: Construct, vpc: IVpc) {

        const role = new Role(
            scope,
            'Role-ec2',
            { assumedBy: new ServicePrincipal('ec2.amazonaws.com') }
        )

        const securityGroup = new SecurityGroup(
            scope,
            'ec2-input-system-sg',
            {
                vpc,
                allowAllOutbound: true,
            }
        )

        securityGroup.addIngressRule(
            Peer.anyIpv4(),
            Port.tcp(22),
            'Allows SSH access from Internet'
        )

        const userData = UserData.forLinux()
        userData.addCommands(
            'sudo yum update',
            'git clone https://github.com/DonaldArmandoArteaga/measurements-admin-system.git',
            'cd measurements-admin-system/src/ec2/',
            'go build .',
            'go run input-system'
        )

        const instance = new Instance(scope, 'ec2-imput-system-instance-1', {
            vpc,
            role: role,
            securityGroup: securityGroup,
            instanceName: 'InputSystemEC2',
            instanceType: InstanceType.of(
                InstanceClass.T2,
                InstanceSize.MICRO
            ),
            machineImage: MachineImage.latestAmazonLinux({
                generation: AmazonLinuxGeneration.AMAZON_LINUX_2,
            }),
            userData,
            userDataCausesReplacement: true,
            keyName: 'ec2-input-system-instance-1-key',
        })


        new CfnOutput(scope, 'InputSystemEC2-output', {
            value: instance.instancePublicIp
        })
    }
}


