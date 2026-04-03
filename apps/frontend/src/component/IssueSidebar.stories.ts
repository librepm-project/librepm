import type { Meta, StoryObj } from '@storybook/vue3';
import IssueSidebar from './IssueSidebar.vue';
import { useIssueStore } from '@/store/issue.store';
import { useProjectStore } from '@/store/project.store';
import { useUserStore } from '@/store/user.store';
import { usePriorityStore } from '@/store/priority.store';
import { useProjectReleaseStore } from '@/store/project-release.store';
import { useProjectReleaseIssueStore } from '@/store/project-release-issue.store';

const meta: Meta<typeof IssueSidebar> = {
  component: IssueSidebar,
  title: 'Components/IssueSidebar',
};
export default meta;

type Story = StoryObj<typeof IssueSidebar>;

const mockIssue = {
  id: 'i1',
  key: 'WED-1',
  summary: 'Book Venue',
  description: '<p>Find and book a suitable venue for the wedding.</p>',
  tracker: { id: 't1', name: 'Task', color: '#6C63FF' },
  trackerId: 't1',
  status: { id: 's1', name: 'In Progress', color: '#FFC107' },
  statusId: 's1',
  priority: { id: 'pr1', name: 'High', color: '#F44336' },
  priorityId: 'pr1',
  project: { id: 'p1', name: 'Wedding Project', codeName: 'WED' },
  projectId: 'p1',
  assignedUser: { id: 'u1', firstName: 'John', lastName: 'Doe', email: 'john@example.com' },
  assignedUserId: 'u1',
  reporterUser: { id: 'u2', firstName: 'Jane', lastName: 'Smith', email: 'jane@example.com' },
};

function seedStores() {
  useIssueStore().current = mockIssue as any;
  useProjectStore().index = [
    { id: 'p1', name: 'Wedding Project', codeName: 'WED', defaultStatusId: 's1', defaultTrackerId: 't1' },
  ] as any;
  useUserStore().index = [
    { id: 'u1', firstName: 'John', lastName: 'Doe', email: 'john@example.com' },
    { id: 'u2', firstName: 'Jane', lastName: 'Smith', email: 'jane@example.com' },
  ] as any;
  usePriorityStore().index = [
    { id: 'pr1', name: 'High', color: '#F44336' },
    { id: 'pr2', name: 'Normal', color: '#FFC107' },
    { id: 'pr3', name: 'Low', color: '#4CAF50' },
  ] as any;
  useProjectReleaseStore().index = [] as any;
  useProjectReleaseIssueStore().index = [] as any;
}

export const Default: Story = {
  render: () => ({
    components: { IssueSidebar },
    setup() {
      seedStores();
      return {};
    },
    template: `<div style="width: 320px;"><IssueSidebar /></div>`,
  }),
};

export const NoIssue: Story = {
  render: () => ({
    components: { IssueSidebar },
    setup() {
      useIssueStore().current = null;
      return {};
    },
    template: `<div style="width: 320px;"><IssueSidebar /></div>`,
  }),
};
