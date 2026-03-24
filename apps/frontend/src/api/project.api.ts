import { createCrudApi } from './baseApi';
import { Project } from '@/lib/interfaces/project.interface';

export default createCrudApi<Project>('project');

