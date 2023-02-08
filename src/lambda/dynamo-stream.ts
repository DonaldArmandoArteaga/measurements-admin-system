import { APIGatewayProxyResult, DynamoDBRecord, DynamoDBStreamEvent } from "aws-lambda";
import { MeasurerFromDynamo, MeasurerMetadataFromDynamo, MeasurerValuesFromDynamo } from "./models/measurers";


export const Handler = async (event: DynamoDBStreamEvent) => {

    const measurer: MeasurerFromDynamo[] = event.Records.map((dynamoDBRecord: DynamoDBRecord): MeasurerFromDynamo => {
        const data = dynamoDBRecord.dynamodb?.NewImage!
        return {
            id: data.id.S!,
            serial: data.serial.S!,
            date: data.date.S!,
            metadata: JSON.parse(data.metadata.S!) as MeasurerMetadataFromDynamo,
            values: JSON.parse(data.values.S!) as MeasurerValuesFromDynamo
        }

    })

    console.log(`Records: ${JSON.stringify(measurer, null, 2)}`);
    return

};