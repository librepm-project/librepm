import type { Meta, StoryObj } from '@storybook/vue3';
import IssueTable from './IssueTable.vue';

const meta: Meta<typeof IssueTable> = {
  component: IssueTable,
  title: 'Components/IssueTable',
};
export default meta;

type Story = StoryObj<typeof IssueTable>;

const items = [
  {
    id: '1', key: 'WED-1', summary: 'Book Venue', description: 'Find and book a venue',
    tracker: { id: '1', name: 'Task', color: '#6C63FF' },
    status: { id: '1', name: 'In Progress', color: '#FFC107' },
    priority: { id: '1', name: 'High', color: '#F44336' },
    assignedUser: { id: 'u1', firstName: 'John', lastName: 'Doe' },
  },
  {
    id: '2', key: 'WED-2', summary: 'Choose Menu', description: 'Select the dinner menu',
    tracker: { id: '2', name: 'Bug', color: '#FF4B4B' },
    status: { id: '2', name: 'Open', color: '#2196F3' },
    priority: { id: '2', name: 'Normal', color: '#FFC107' },
    assignedUser: null,
  },
  {
    id: '3', key: 'WED-3', summary: 'Send Invitations', description: 'Mail all invitations',
    tracker: { id: '1', name: 'Task', color: '#6C63FF' },
    status: { id: '3', name: 'Done', color: '#4CAF50' },
    priority: { id: '3', name: 'Low', color: '#4CAF50' },
    assignedUser: { id: 'u2', firstName: 'Jane', lastName: 'Smith' },
  },
];

export const Default: Story = {
  args: { items },
};

export const KeyAndSummaryOnly: Story = {
  args: {
    items,
    columns: ['key', 'summary'],
  },
};

export const WithGroupBy: Story = {
  args: {
    items,
    groupBy: 'status',
  },
};

export const Empty: Story = {
  args: { items: [] },
};
