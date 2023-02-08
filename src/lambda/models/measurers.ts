export interface MeasurerFromDynamo {
    id: string
    serial: string
    date: string
    metadata: MeasurerMetadataFromDynamo
    values: MeasurerValuesFromDynamo
}

export interface MeasurerMetadataFromDynamo {
    DateInserted: string
}

export interface MeasurerValuesFromDynamo {
    temperature?: number
}

