import { CfnOutput } from "aws-cdk-lib";
import { AttributeType, BillingMode, StreamViewType, Table } from "aws-cdk-lib/aws-dynamodb";
import { Construct } from "constructs";

export class InputSystemStackDynamoDB {
    private readonly inputSystemTable: Table
    constructor(scope: Construct) {

        this.inputSystemTable = new Table(scope, 'InputSystemDynamoTable', {
            partitionKey: { name: 'id', type: AttributeType.STRING },
            sortKey: { name: 'serial', type: AttributeType.STRING },
            billingMode: BillingMode.PAY_PER_REQUEST,
            stream: StreamViewType.NEW_IMAGE
        });

        new CfnOutput(scope, 'InputSystemDynamo-output', {
            value: this.inputSystemTable.tableName
        })

    }

    get GetInputSystemTable(): Table {
        return this.inputSystemTable
    }
}