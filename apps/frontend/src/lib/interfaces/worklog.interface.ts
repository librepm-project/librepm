import { User } from './user.interface';

export interface Worklog {
  id?: string;
  seconds: number;
  comment?: string;
  loggedAt: string;
  user?: User;
}
