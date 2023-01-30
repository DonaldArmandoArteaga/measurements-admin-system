import { CfnOutput, Tags, Token } from "aws-cdk-lib"
import { SecurityGroup, Peer, Port, UserData, Instance, InstanceType, InstanceClass, InstanceSize, MachineImage, AmazonLinuxGeneration, CloudFormationInit, InitFile, InitConfig, IVpc, InitCommand } from "aws-cdk-lib/aws-ec2"
import { Effect, Policy, PolicyStatement, Role, ServicePrincipal } from "aws-cdk-lib/aws-iam"
import { Construct } from "constructs"
import * as fs from 'fs'
import * as path from 'path';
import { InputSystemStackQueue } from "./sqs"


export class InputSystemStackEC2 {
    constructor(scope: Construct, vpc: IVpc, tableName: string, inputSystemStackQueue: InputSystemStackQueue) {

        const role = new Role(
            scope,
            'Role-ec2',
            { assumedBy: new ServicePrincipal('ec2.amazonaws.com') }
        )

        role.attachInlinePolicy(new Policy(scope, "RetreiveMessagePolicy", {
            statements: [
                new PolicyStatement({
                    actions: [
                        "sqs:*",
                    ],
                    effect: Effect.ALLOW,
                    resources: [inputSystemStackQueue.getQueueARN],
                }),
            ],
        }))

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
        userData.addCommands('sudo su', `export QUEUE_NAME=${inputSystemStackQueue.getQueueName}`, `export DYNAMO_TABLE_NAME=${tableName}`)
        userData.addCommands(fs.readFileSync(path.join(__dirname, `../../src/ec2/scripts/run.sh`), 'utf8'))

        const instance = new Instance(scope, 'ec2-imput-system-instance-1', {
            vpc,
            role,
            securityGroup: securityGroup,
            instanceType: InstanceType.of(
                InstanceClass.T2,
                InstanceSize.MICRO
            ),
            machineImage: MachineImage.latestAmazonLinux({
                generation: AmazonLinuxGeneration.AMAZON_LINUX_2,
            }),
            userData,
            userDataCausesReplacement: true,
            keyName: 'ec2-input-system-instance-1-key'
        })



        new CfnOutput(scope, 'InputSystemEC2-output', {
            value: instance.instancePublicIp
        })

    }
}


