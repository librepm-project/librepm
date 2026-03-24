import { createCrudApi } from './baseApi';
import { User } from '@/lib/interfaces/user.interface';

export default createCrudApi<User>('user');

