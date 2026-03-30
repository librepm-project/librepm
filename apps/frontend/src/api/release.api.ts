import { createCrudApi } from './baseApi';
import { Release } from '@/lib/interfaces/release.interface';

export default createCrudApi<Release>('release');
