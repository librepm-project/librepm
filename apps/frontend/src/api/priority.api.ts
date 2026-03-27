import { createCrudApi } from './baseApi';
import { Priority } from '@/lib/interfaces/priority.interface';

export default createCrudApi<Priority>('priority');
