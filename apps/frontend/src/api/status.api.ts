import { createCrudApi } from './baseApi';
import { Status } from '@/lib/interfaces/status.interface';

export default createCrudApi<Status>('status');
