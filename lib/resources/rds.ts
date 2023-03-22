import { InstanceClass, InstanceSize, InstanceType, IVpc, Peer, Port, SecurityGroup, SubnetType } from "aws-cdk-lib/aws-ec2";
import { Construct } from "constructs";
import { Function } from "aws-cdk-lib/aws-lambda";
import { DatabaseInstance, DatabaseInstanceEngine, PostgresEngineVersion, Credentials } from "aws-cdk-lib/aws-rds";
import { CfnOutput, SecretValue } from "aws-cdk-lib";

export class InputSystemStackRDS {

    constructor(scope: Construct, vpc: IVpc, lambda: Function) {

        const securityGroup = new SecurityGroup(
            scope,
            'rds-input-system-sg',
            {
                vpc,
                allowAllOutbound: true,
            }
        )

        securityGroup.addIngressRule(
            Peer.anyIpv4(),
            Port.tcp(5432),
            'Allows Database access from Internet'
        )

        const instance = new DatabaseInstance(scope, 'Instance', {
            engine: DatabaseInstanceEngine.postgres({ version: PostgresEngineVersion.VER_14_2 }),
            instanceType: InstanceType.of(
                InstanceClass.T3,
                InstanceSize.MICRO
            ),
            credentials: Credentials.fromPassword('postgres', SecretValue.unsafePlainText("Armando20109210*")),
            vpc,
            vpcSubnets: {
                subnetType: SubnetType.PUBLIC,
            },
            publiclyAccessible: true,
            securityGroups: [securityGroup],

        });


        new CfnOutput(scope, 'InputSystemRDS-output', {
            value: `Address: ${instance.dbInstanceEndpointAddress}, port: ${instance.dbInstanceEndpointPort}`
        })


    }
}