import { Release } from './release.interface';
import { Project } from './project.interface';

export interface ProjectRelease {
  id?: string;
  release_id: string;
  project_id: string;
  release?: Release;
  project?: Project;
}
