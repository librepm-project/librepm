export interface FilterCondition {
    field: string;
    op: string;
    value: string;
};

export interface Filter {
    name: string;
    description: string;
    conditions: FilterCondition[];
};
