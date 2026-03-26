import { get, put } from './api';
import { Setting } from '@/lib/interfaces/setting.interface';

export default {
    index: () => get<Setting[]>('setting'),
    update: (key: string, value: string) => put<void>(`setting/${key}`, { value }),
};
