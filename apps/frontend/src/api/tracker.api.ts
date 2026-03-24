import { createCrudApi } from './baseApi';
import { Tracker } from '@/lib/interfaces/tracker.interface';

export default createCrudApi<Tracker>('tracker');
