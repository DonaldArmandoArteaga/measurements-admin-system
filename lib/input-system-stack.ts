import { Stack, StackProps } from "aws-cdk-lib";
import { Construct } from "constructs";
import { InputSystemStackAPIGateway } from "./resources/api-gateway";
import { InputSystemStackCommonInstances } from "./resources/common/intances";
import { InputSystemStackDynamoDB } from "./resources/dynamodb";
import { InputSystemStackEC2 } from "./resources/ec2";
import { InputSystemStackLambda } from "./resources/lambda";
import { InputSystemStackRDS } from "./resources/rds";
import { InputSystemStackQueue } from "./resources/sqs";
import { InputSystemStackParameterStore } from "./resources/parameter-store";

export class InputSystemStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    const inputSystemStackCommonInstances = new InputSystemStackCommonInstances(
      this
    );

    const inputSystemStackQueue = new InputSystemStackQueue(this);

    const apiGateway = new InputSystemStackAPIGateway(
      this,
      inputSystemStackQueue
    );
    const inputSystemStackDynamoDB = new InputSystemStackDynamoDB(this);
    new InputSystemStackParameterStore(
      this,
      inputSystemStackQueue,
      inputSystemStackDynamoDB.GetInputSystemTable,
      apiGateway.ApiURL
    );
    new InputSystemStackEC2(
      this,
      inputSystemStackCommonInstances.GeDefaultVPC,
      inputSystemStackDynamoDB.GetInputSystemTable,
      inputSystemStackQueue,
      apiGateway.ApiURL
    );
    const inputSystemStackLambda = new InputSystemStackLambda(
      this,
      inputSystemStackCommonInstances.GeDefaultVPC,
      inputSystemStackDynamoDB.GetInputSystemTable
    );
    //new InputSystemStackRDS(this, inputSystemStackCommonInstances.GeDefaultVPC, inputSystemStackLambda.GetDynamoStream)
  }
}
