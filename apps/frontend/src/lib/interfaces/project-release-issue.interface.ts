import { Issue } from './issue.interface';

export interface ProjectReleaseIssue {
  id?: string;
  project_release_id: string;
  issue_id: string;
  issue?: Issue;
}
