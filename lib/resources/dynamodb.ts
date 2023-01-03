import { AttributeType, BillingMode, StreamViewType, Table } from "aws-cdk-lib/aws-dynamodb";
import { Construct } from "constructs";

export class InputSystemStackDynamoDB {
    private readonly inputSystemTable: Table
    constructor(scope: Construct) {

        this.inputSystemTable = new Table(scope, 'InputSystemDynamoTable', {
            tableName:'raw-measurements-table',
            partitionKey: { name: 'serial', type: AttributeType.STRING },
            sortKey: { name: 'date', type: AttributeType.STRING },
            billingMode: BillingMode.PAY_PER_REQUEST,
            stream: StreamViewType.NEW_IMAGE
        });

    }

    get GetInputSystemTable(): Table {
        return this.inputSystemTable
    }
}