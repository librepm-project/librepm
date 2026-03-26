import { User } from './user.interface';

export interface Comment {
  id?: string;
  content: string;
  entityType: string;
  entityId: string;
  user?: User;
  createdAt?: string;
  updatedAt?: string;
}
