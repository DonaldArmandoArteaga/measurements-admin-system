import {
  ParameterTier,
  ParameterType,
  StringParameter,
} from "aws-cdk-lib/aws-ssm";
import { Construct } from "constructs";
import { InputSystemStackQueue } from "./sqs";
import { Table } from "aws-cdk-lib/aws-dynamodb";
import { Stack } from "aws-cdk-lib";

export class InputSystemStackParameterStore {
  constructor(
    scope: Construct,
    queue: InputSystemStackQueue,
    table: Table,
    apiURL: string
  ) {
    new StringParameter(scope, "QUEUE_NAME_PARAMETER", {
      parameterName: "QUEUE_NAME",
      stringValue: queue.getQueueName,
      description: "QUEUE_NAME_PARAMETER",
      type: ParameterType.STRING,
      tier: ParameterTier.STANDARD,
    });

    new StringParameter(scope, "DYNAMO_TABLE_NAME_PARAMETER", {
      parameterName: "DYNAMO_TABLE_NAME",
      stringValue: table.tableName,
      description: "DYNAMO_TABLE_NAME_PARAMETER",
      type: ParameterType.STRING,
      tier: ParameterTier.STANDARD,
    });

    new StringParameter(scope, "REGION_PARAMETER", {
      parameterName: "REGION",
      stringValue: Stack.of(scope).region,
      description: "REGION_PARAMETER",
      type: ParameterType.STRING,
      tier: ParameterTier.STANDARD,
    });

    new StringParameter(scope, "API_URL_PARAMETER", {
      parameterName: "API_URL",
      stringValue: apiURL,
      description: "API_URL_PARAMETER",
      type: ParameterType.STRING,
      tier: ParameterTier.STANDARD,
    });
  }
}
