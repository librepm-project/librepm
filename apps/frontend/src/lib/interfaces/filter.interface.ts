export interface FilterCondition {
    id?: string;
    field: string;
    op: string;
    value: string;
}

export interface Filter {
    id?: string;
    name: string;
    description?: string;
    conditions?: FilterCondition[];
}

export interface FilterConditionOperator {
    key: string;
    label: string;
}

export interface FilterConditionField {
    key: string;
    label: string;
    valueType: 'text' | 'select';
    valueEndpoint?: string;
    operators: FilterConditionOperator[];
}

export interface FilterConditionOptions {
    fields: FilterConditionField[];
}
