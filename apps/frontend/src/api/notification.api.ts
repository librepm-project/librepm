import { Notification } from '@/lib/interfaces/notification.interface';
import api from '@/api/api';

const getUnread = async (): Promise<Notification[]> => {
  const response = await api.apiCall().get('/notification');
  return response.data as Notification[];
};

const markAsRead = async (id: string): Promise<void> => {
  await api.apiCall().patch(`/notification/${id}/read`);
};

export default {
  getUnread,
  markAsRead,
};
