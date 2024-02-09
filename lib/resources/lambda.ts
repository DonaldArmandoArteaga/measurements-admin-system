import { Duration } from "aws-cdk-lib";
import { Table } from "aws-cdk-lib/aws-dynamodb";
import { IVpc, SecurityGroup, SubnetType } from "aws-cdk-lib/aws-ec2";
import { Runtime, StartingPosition } from "aws-cdk-lib/aws-lambda";
import { DynamoEventSource } from "aws-cdk-lib/aws-lambda-event-sources";
import { NodejsFunction } from "aws-cdk-lib/aws-lambda-nodejs";
import { Construct } from "constructs";
import * as path from 'path';
export class InputSystemStackLambda {
    private readonly DynamoStream: NodejsFunction
    constructor(scope: Construct, vpc: IVpc, inputSystemTable: Table) {

        this.DynamoStream = new NodejsFunction(scope, 'dynamo-stream-lambda', {
            memorySize: 128,
            timeout: Duration.seconds(5),
            runtime: Runtime.NODEJS_16_X,
            handler: 'Handler',
            entry: path.join(__dirname, `../../src/lambda/dynamo-stream.ts`),
            //vpc,
        });

        this.DynamoStream.addEventSource(new DynamoEventSource(inputSystemTable, {
            startingPosition: StartingPosition.LATEST,
            batchSize: 20
        }));

    }

    get GetDynamoStream(): NodejsFunction {
        return this.DynamoStream
    }


}