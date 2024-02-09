import { Aws, CfnOutput } from "aws-cdk-lib";
import { AwsIntegration, RestApi } from "aws-cdk-lib/aws-apigateway";
import { Role, ServicePrincipal, Policy, PolicyStatement, Effect } from "aws-cdk-lib/aws-iam";
import { Construct } from "constructs";
import { InputSystemStackQueue } from "./sqs";

export class InputSystemStackAPIGateway {
    private readonly apiV1: RestApi
    constructor(scope: Construct, inputSystemStackQueue: InputSystemStackQueue) {

        const credentialsRole = new Role(scope, "Role-apigateway", {
            assumedBy: new ServicePrincipal("apigateway.amazonaws.com"),
        });

        credentialsRole.attachInlinePolicy(
            new Policy(scope, "SendMessagePolicy", {
                statements: [
                    new PolicyStatement({
                        actions: ["sqs:SendMessage"],
                        effect: Effect.ALLOW,
                        resources: [inputSystemStackQueue.getQueueARN],
                    }),
                ],
            })
        );

        const inputGatewayIntegratedInputQueue = new AwsIntegration({
            service: 'sqs',
            path: `${Aws.ACCOUNT_ID}/${inputSystemStackQueue.getQueueName}`,
            integrationHttpMethod: "POST",
            options: {
                credentialsRole,
                requestParameters: {
                    "integration.request.header.Content-Type": `'application/x-www-form-urlencoded'`,
                },
                requestTemplates: {
                    "application/json": `Action=SendMessage&MessageBody=$input.body`,
                },
                integrationResponses: [
                    {
                        statusCode: "202",
                        responseTemplates: {
                            "text/plain": `received`,
                        },
                    },
                ],
            },
        });

        this.apiV1 = new RestApi(scope, 'InputSystemGateway');
        const v1 = this.apiV1.root.addResource('v1');
        const books = v1.addResource('input-data');
        books.addMethod('POST', inputGatewayIntegratedInputQueue, { methodResponses: [{ statusCode: "202" }] });


        new CfnOutput(scope, 'InputSystemApiGateway-output', {
            value: this.apiV1.url
        })
    }


    get ApiURL(): string {
        return this.apiV1.url
    }
}