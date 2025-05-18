export interface ProjectIssueProperty {
  trackers: {
    id: string;
    name: string;
    statuses: {
      id: string;
      name: string;
    }[];
  }[];
}
