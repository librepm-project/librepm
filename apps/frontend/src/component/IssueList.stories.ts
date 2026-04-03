import type { Meta, StoryObj } from '@storybook/vue3';
import IssueList from './IssueList.vue';

const meta: Meta<typeof IssueList> = {
  component: IssueList,
  title: 'Components/IssueList',
};
export default meta;

type Story = StoryObj<typeof IssueList>;

const items = [
  {
    id: 'i1', key: 'WED-1', summary: 'Book Venue', description: 'Find and book a venue',
    tracker: { id: 't1', name: 'Task', color: '#6C63FF' },
    status: { id: 's1', name: 'In Progress', color: '#FFC107' },
    priority: { id: 'pr1', name: 'High', color: '#F44336' },
    assignedUser: { id: 'u1', firstName: 'John', lastName: 'Doe' },
  },
  {
    id: 'i2', key: 'WED-2', summary: 'Choose Menu', description: 'Select the dinner menu',
    tracker: { id: 't2', name: 'Bug', color: '#FF4B4B' },
    status: { id: 's2', name: 'Open', color: '#2196F3' },
    priority: { id: 'pr2', name: 'Normal', color: '#FFC107' },
    assignedUser: null,
  },
  {
    id: 'i3', key: 'WED-3', summary: 'Send Invitations', description: 'Mail all invitations',
    tracker: { id: 't1', name: 'Task', color: '#6C63FF' },
    status: { id: 's3', name: 'Done', color: '#4CAF50' },
    priority: { id: 'pr3', name: 'Low', color: '#4CAF50' },
    assignedUser: { id: 'u2', firstName: 'Jane', lastName: 'Smith' },
  },
];

export const Default: Story = {
  args: { items },
};

export const WithSettingsToggle: Story = {
  args: {
    items,
    showSettingsToggle: true,
  },
};

export const Empty: Story = {
  args: { items: [] },
};
