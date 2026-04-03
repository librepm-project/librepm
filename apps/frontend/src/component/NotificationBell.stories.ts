import type { Meta, StoryObj } from '@storybook/vue3';
import NotificationBell from './NotificationBell.vue';
import { useNotificationStore } from '@/store/notification.store';

const meta: Meta<typeof NotificationBell> = {
  component: NotificationBell,
  title: 'Components/NotificationBell',
};
export default meta;

type Story = StoryObj<typeof NotificationBell>;

const notifications = [
  { id: 'n1', type: 'issue_assigned', entityType: 'issue', entityId: 'i1', read: false, createdAt: '2025-03-01T10:00:00Z' },
  { id: 'n2', type: 'comment_added', entityType: 'issue', entityId: 'i2', read: false, createdAt: '2025-03-01T09:00:00Z' },
  { id: 'n3', type: 'mention', entityType: 'issue', entityId: 'i3', read: false, createdAt: '2025-03-01T08:00:00Z' },
];

export const WithNotifications: Story = {
  render: () => ({
    components: { NotificationBell },
    setup() {
      const store = useNotificationStore();
      store.items = notifications as any;
      return {};
    },
    template: `<div style="padding: 16px; display: flex; justify-content: flex-end;"><NotificationBell /></div>`,
  }),
};

export const NoNotifications: Story = {
  render: () => ({
    components: { NotificationBell },
    setup() {
      const store = useNotificationStore();
      store.items = [];
      return {};
    },
    template: `<div style="padding: 16px; display: flex; justify-content: flex-end;"><NotificationBell /></div>`,
  }),
};
