import type { Meta, StoryObj } from '@storybook/vue3';
import IssueForm from './IssueForm.vue';
import { useProjectStore } from '@/store/project.store';
import { useUserStore } from '@/store/user.store';
import { usePriorityStore } from '@/store/priority.store';

const meta: Meta<typeof IssueForm> = {
  component: IssueForm,
  title: 'Components/IssueForm',
};
export default meta;

type Story = StoryObj<typeof IssueForm>;

const projects = [
  { id: 'p1', name: 'Wedding Project', codeName: 'WED', defaultStatusId: 's1', defaultTrackerId: 't1' },
  { id: 'p2', name: 'Startup Launch', codeName: 'START', defaultStatusId: null, defaultTrackerId: null },
];

const users = [
  { id: 'u1', email: 'john@example.com', firstName: 'John', lastName: 'Doe' },
  { id: 'u2', email: 'jane@example.com', firstName: 'Jane', lastName: 'Smith' },
];

const priorities = [
  { id: 'pr1', name: 'High', color: '#F44336' },
  { id: 'pr2', name: 'Normal', color: '#FFC107' },
  { id: 'pr3', name: 'Low', color: '#4CAF50' },
];

export const Create: Story = {
  render: (args) => ({
    components: { IssueForm },
    setup() {
      const projectStore = useProjectStore();
      const userStore = useUserStore();
      const priorityStore = usePriorityStore();
      projectStore.index = projects as any;
      userStore.index = users as any;
      priorityStore.index = priorities as any;
      return { args };
    },
    template: `<IssueForm v-bind="args" />`,
  }),
  args: {
    submitButtonText: 'global.create',
    onSubmit: (item: any) => console.log('submit', item),
  },
};

export const Edit: Story = {
  render: (args) => ({
    components: { IssueForm },
    setup() {
      const projectStore = useProjectStore();
      const userStore = useUserStore();
      const priorityStore = usePriorityStore();
      projectStore.index = projects as any;
      userStore.index = users as any;
      priorityStore.index = priorities as any;
      return { args };
    },
    template: `<IssueForm v-bind="args" />`,
  }),
  args: {
    submitButtonText: 'global.save',
    onSubmit: (item: any) => console.log('submit', item),
    initialData: {
      id: 'i1',
      key: 'WED-1',
      summary: 'Book Venue',
      description: '<p>Find and book a suitable venue for the wedding.</p>',
      projectId: 'p1',
      priorityId: 'pr1',
      assignedUserId: 'u1',
    },
  },
};
