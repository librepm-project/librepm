import type { Meta, StoryObj } from '@storybook/vue3';
import CommentPanel from './CommentPanel.vue';

const meta: Meta<typeof CommentPanel> = {
  component: CommentPanel,
  title: 'Components/CommentPanel',
};
export default meta;

type Story = StoryObj<typeof CommentPanel>;

const user = { id: 'u1', firstName: 'John', lastName: 'Doe', email: 'john@example.com' };

const comments = [
  {
    id: 'c1',
    content: '<p>This venue looks great! I already contacted them.</p>',
    entityType: 'issue',
    entityId: 'i1',
    user,
    createdAt: '2025-03-01T10:00:00Z',
    updatedAt: '2025-03-01T10:00:00Z',
  },
  {
    id: 'c2',
    content: '<p>The pricing is a bit high. Let me check alternatives.</p>',
    entityType: 'issue',
    entityId: 'i1',
    user: { id: 'u2', firstName: 'Jane', lastName: 'Smith', email: 'jane@example.com' },
    createdAt: '2025-03-01T11:30:00Z',
    updatedAt: '2025-03-01T11:30:00Z',
  },
];

export const Default: Story = {
  args: {
    issueId: 'i1',
    comments,
  },
};

export const Empty: Story = {
  args: {
    issueId: 'i1',
    comments: [],
  },
};
