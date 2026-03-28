import { createCrudApi } from './baseApi';
import { Project } from '@/lib/interfaces/project.interface';
import { Tracker } from '@/lib/interfaces/tracker.interface';
import { Status } from '@/lib/interfaces/status.interface';
import api from '@/api/api';

const crudApi = createCrudApi<Project>('project');

const trackers = async (projectId: string): Promise<Tracker[]> => {
    const response = await api.apiCall().get(`/project/${projectId}/tracker`);
    return response.data as Tracker[];
}

const statuses = async (projectId: string): Promise<Status[]> => {
    const response = await api.apiCall().get(`/project/${projectId}/status`);
    return response.data as Status[];
}

export default {
    ...crudApi,
    trackers,
    statuses,
}
