import type { Meta, StoryObj } from '@storybook/vue3';
import IssueAuditLogPanel from './IssueAuditLogPanel.vue';

const meta: Meta<typeof IssueAuditLogPanel> = {
  component: IssueAuditLogPanel,
  title: 'Components/IssueAuditLogPanel',
};
export default meta;

type Story = StoryObj<typeof IssueAuditLogPanel>;

const user = { id: 'u1', firstName: 'John', lastName: 'Doe', email: 'john@example.com' };

const auditLogs = [
  {
    id: 'al1',
    eventType: 'field_changed',
    field: 'status',
    oldValue: 'Open',
    newValue: 'In Progress',
    meta: '',
    user,
    createdAt: '2025-03-01T10:00:00Z',
  },
  {
    id: 'al2',
    eventType: 'field_changed',
    field: 'assignee',
    oldValue: '',
    newValue: 'John Doe',
    meta: '',
    user,
    createdAt: '2025-03-01T09:30:00Z',
  },
  {
    id: 'al3',
    eventType: 'created',
    field: '',
    oldValue: '',
    newValue: '',
    meta: '',
    user,
    createdAt: '2025-03-01T09:00:00Z',
  },
];

export const Default: Story = {
  args: {
    issueId: 'i1',
    auditLogs,
  },
};

export const Empty: Story = {
  args: {
    issueId: 'i1',
    auditLogs: [],
  },
};
