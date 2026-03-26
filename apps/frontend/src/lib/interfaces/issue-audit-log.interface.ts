import { User } from './user.interface';

export interface IssueAuditLog {
  id: string;
  eventType: string;
  field: string;
  oldValue: string;
  newValue: string;
  meta: string;
  user: User;
  createdAt: string;
}
