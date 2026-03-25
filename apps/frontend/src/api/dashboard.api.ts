import { Dashboard } from '@/lib/interfaces/dashboard.interface';
import { createCrudApi } from './baseApi';

export default createCrudApi<Dashboard>('dashboard');
