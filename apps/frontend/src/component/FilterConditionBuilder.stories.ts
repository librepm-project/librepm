import type { Meta, StoryObj } from '@storybook/vue3';
import FilterConditionBuilder from './FilterConditionBuilder.vue';
import { useProjectStore } from '@/store/project.store';
import { useTrackerStore } from '@/store/tracker.store';
import { useStatusStore } from '@/store/status.store';
import { useUserStore } from '@/store/user.store';
import { usePriorityStore } from '@/store/priority.store';

const meta: Meta<typeof FilterConditionBuilder> = {
  component: FilterConditionBuilder,
  title: 'Components/FilterConditionBuilder',
};
export default meta;

type Story = StoryObj<typeof FilterConditionBuilder>;

function seedStores() {
  useProjectStore().index = [
    { id: 'p1', name: 'Wedding Project', codeName: 'WED' },
    { id: 'p2', name: 'Startup Launch', codeName: 'START' },
  ] as any;
  useTrackerStore().index = [
    { id: 't1', name: 'Bug', color: '#FF4B4B' },
    { id: 't2', name: 'Task', color: '#6C63FF' },
  ] as any;
  useStatusStore().index = [
    { id: 's1', name: 'Open', color: '#2196F3' },
    { id: 's2', name: 'In Progress', color: '#FFC107' },
    { id: 's3', name: 'Done', color: '#4CAF50' },
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
}

export const Empty: Story = {
  render: (args) => ({
    components: { FilterConditionBuilder },
    setup() {
      seedStores();
      return { args };
    },
    template: `<FilterConditionBuilder v-bind="args" />`,
  }),
  args: {
    modelValue: [],
  },
};

export const WithConditions: Story = {
  render: (args) => ({
    components: { FilterConditionBuilder },
    setup() {
      seedStores();
      return { args };
    },
    template: `<FilterConditionBuilder v-bind="args" />`,
  }),
  args: {
    modelValue: [
      { id: 'c1', field: 'summary', op: 'contains', value: 'venue' },
      { id: 'c2', field: 'status', op: 'eq', value: 's1' },
    ],
  },
};
