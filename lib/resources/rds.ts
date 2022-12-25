
import { Duration, RemovalPolicy, CfnOutput } from "aws-cdk-lib";
import { SubnetType, InstanceType, InstanceClass, InstanceSize, Port, IVpc, SecurityGroup } from "aws-cdk-lib/aws-ec2";
import { DatabaseInstance, DatabaseInstanceEngine, PostgresEngineVersion, Credentials } from "aws-cdk-lib/aws-rds";
import { Construct } from "constructs";
import { Function } from "aws-cdk-lib/aws-lambda";


export class InputSystemStackRDS {

    constructor(scope: Construct, vpc: IVpc, lambda: Function) {

        // const dbSecurityGroup = new SecurityGroup(scope, 'InputSystemPostgresRDS-securityGroup', {
        //     vpc,
        // });

        // const dbInstance = new DatabaseInstance(scope, 'InputSystemPostgresRDS', {
        //     vpc,
        //     vpcSubnets: vpc.selectSubnets({
        //         subnetType: SubnetType.PRIVATE_ISOLATED,
        //     }),
        //     engine: DatabaseInstanceEngine.postgres({
        //         version: PostgresEngineVersion.VER_14,
        //     }),
        //     instanceType: InstanceType.of(
        //         InstanceClass.BURSTABLE3,
        //         InstanceSize.SMALL,
        //     ),
        //     credentials: Credentials.fromGeneratedSecret('postgres'),
        //     multiAz: false,
        //     allocatedStorage: 100,
        //     maxAllocatedStorage: 100,
        //     allowMajorVersionUpgrade: false,
        //     autoMinorVersionUpgrade: true,
        //     backupRetention: Duration.days(0),
        //     deleteAutomatedBackups: true,
        //     removalPolicy: RemovalPolicy.DESTROY,
        //     deletionProtection: false,
        //     databaseName: 'measurements',
        //     publiclyAccessible: false,
        //     securityGroups: [dbSecurityGroup],
        // });

        // dbInstance.connections.allowFrom(lambda, Port.tcp(5432));

        // dbInstance.secret?.grantRead(lambda);

        // new CfnOutput(scope, 'InputSystemPostgresRDS-hostName', {
        //     value: dbInstance.instanceEndpoint.hostname,
        // });

        // new CfnOutput(scope, 'InputSystemPostgresRDS-secretName', {
        //     value: dbInstance.secret?.secretName!,
        // });
    }
}