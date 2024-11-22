export interface FilterCondition {
    id: string;
    field: string;
    op: string;
    value: string;
};

export interface Filter {
    id: string;
    name: string;
    description: string;
    conditions: FilterCondition[];
};
