import type { Meta, StoryObj } from '@storybook/vue3';
import RelatedIssuesPanel from './RelatedIssuesPanel.vue';
import { useIssueStore } from '@/store/issue.store';

const meta: Meta<typeof RelatedIssuesPanel> = {
  component: RelatedIssuesPanel,
  title: 'Components/RelatedIssuesPanel',
};
export default meta;

type Story = StoryObj<typeof RelatedIssuesPanel>;

const issueA = {
  id: 'i1', key: 'WED-1', summary: 'Book Venue', description: '',
  tracker: { id: 't1', name: 'Task', color: '#6C63FF' },
  status: { id: 's1', name: 'In Progress', color: '#FFC107' },
};

const issueB = {
  id: 'i2', key: 'WED-2', summary: 'Choose Menu', description: '',
  tracker: { id: 't1', name: 'Task', color: '#6C63FF' },
  status: { id: 's2', name: 'Open', color: '#2196F3' },
};

const issueC = {
  id: 'i3', key: 'WED-3', summary: 'Send Invitations', description: '',
  tracker: { id: 't2', name: 'Bug', color: '#FF4B4B' },
  status: { id: 's3', name: 'Done', color: '#4CAF50' },
};

const relations = [
  { id: 'r1', type: 'blocks' as const, issueA, issueB, directionFromCurrent: 'blocks' as const },
  { id: 'r2', type: 'related' as const, issueA: issueA, issueB: issueC, directionFromCurrent: 'related' as const },
];

export const Default: Story = {
  render: (args) => ({
    components: { RelatedIssuesPanel },
    setup() {
      const issueStore = useIssueStore();
      issueStore.index = [issueA, issueB, issueC] as any;
      return { args };
    },
    template: `<RelatedIssuesPanel v-bind="args" />`,
  }),
  args: {
    currentIssueId: 'i1',
    relations,
  },
};

export const Empty: Story = {
  render: (args) => ({
    components: { RelatedIssuesPanel },
    setup() {
      const issueStore = useIssueStore();
      issueStore.index = [issueA, issueB, issueC] as any;
      return { args };
    },
    template: `<RelatedIssuesPanel v-bind="args" />`,
  }),
  args: {
    currentIssueId: 'i1',
    relations: [],
  },
};
