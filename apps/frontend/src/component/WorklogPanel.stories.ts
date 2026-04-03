import type { Meta, StoryObj } from '@storybook/vue3';
import WorklogPanel from './WorklogPanel.vue';

const meta: Meta<typeof WorklogPanel> = {
  component: WorklogPanel,
  title: 'Components/WorklogPanel',
};
export default meta;

type Story = StoryObj<typeof WorklogPanel>;

const user = { id: 'u1', firstName: 'John', lastName: 'Doe', email: 'john@example.com' };

const worklogs = [
  { id: 'w1', seconds: 7200, comment: 'Visited the venue and took notes', loggedAt: '2025-03-01', user },
  { id: 'w2', seconds: 3600, comment: 'Phone call with the coordinator', loggedAt: '2025-03-02', user },
  { id: 'w3', seconds: 1800, comment: '', loggedAt: '2025-03-03', user },
];

export const Default: Story = {
  args: {
    issueId: 'i1',
    worklogs,
  },
};

export const Empty: Story = {
  args: {
    issueId: 'i1',
    worklogs: [],
  },
};
