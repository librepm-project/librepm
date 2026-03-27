import { defineStore } from 'pinia';
import { Notification } from '@/lib/interfaces/notification.interface';
import notificationApi from '@/api/notification.api';

interface NotificationStore {
  items: Notification[];
}

export const useNotificationStore = defineStore('notification', {
  state: (): NotificationStore => ({
    items: [],
  }),
  actions: {
    async fetchUnread() {
      this.items = await notificationApi.getUnread();
    },
    async markAsRead(id: string) {
      await notificationApi.markAsRead(id);
      this.items = this.items.filter((n) => n.id !== id);
    },
    addFromWs(notification: Notification) {
      this.items.unshift(notification);
    },
  },
});
