import { createCrudApi } from './baseApi';
import { Filter } from '@/lib/interfaces/filter.interface';

export default createCrudApi<Filter>('filter');

