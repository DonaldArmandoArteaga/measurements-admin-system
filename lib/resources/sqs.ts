import { Duration } from "aws-cdk-lib";
import { SqsEventSource } from "aws-cdk-lib/aws-lambda-event-sources";
import { Queue } from "aws-cdk-lib/aws-sqs";
import { Construct } from "constructs";

export class InputSystemStackQueue {
    private readonly queue: Queue

    constructor(scope: Construct) {

        const deadLetterQueue = new Queue(scope, 'InputSystemQueue-dlq', {
            retentionPeriod: Duration.days(10),
        });

        this.queue = new Queue(scope, 'InputSystemQueue', {
            visibilityTimeout: Duration.seconds(300),
            deadLetterQueue: {
                queue: deadLetterQueue,
                maxReceiveCount: 1,
            },
        });
    }

    get getQueueName(): string {
        return this.queue.queueName
    }

    get getQueueARN(): string {
        return this.queue.queueArn
    }
}