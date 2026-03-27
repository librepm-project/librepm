export interface Notification {
  id: string;
  type: string;
  entityType?: string;
  entityId?: string;
  read: boolean;
  createdAt: string;
}
