import api from '@/api/api';
import { createCrudApi } from './baseApi';
import { Filter, FilterConditionOptions } from '@/lib/interfaces/filter.interface';

const crudApi = createCrudApi<Filter>('filter');

export default {
    ...crudApi,

    async conditionOptions(): Promise<FilterConditionOptions> {
        const response = await api.apiCall().get('/filter/condition-options');
        return response.data;
    },
};
