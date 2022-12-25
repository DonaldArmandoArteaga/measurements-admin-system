import { Context, APIGatewayProxyResult, DynamoDBStreamEvent } from "aws-lambda";


export const Handler = async (event: DynamoDBStreamEvent): Promise<APIGatewayProxyResult> => {
    console.log(`Records: ${JSON.stringify(event.Records, null, 2)}`);

    return {
        statusCode: 200,
        body: JSON.stringify({
            message: 'hello Donald Armando',
        }),
    };
};