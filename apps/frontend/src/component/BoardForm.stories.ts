import type { Meta, StoryObj } from '@storybook/vue3';
import BoardForm from './BoardForm.vue';
import { useStatusStore } from '@/store/status.store';

const meta: Meta<typeof BoardForm> = {
  component: BoardForm,
  title: 'Components/BoardForm',
};
export default meta;

type Story = StoryObj<typeof BoardForm>;

const statuses = [
  { id: '1', name: 'Open', color: '#2196F3' },
  { id: '2', name: 'In Progress', color: '#FFC107' },
  { id: '3', name: 'Done', color: '#4CAF50' },
];

export const Create: Story = {
  render: (args) => ({
    components: { BoardForm },
    setup() {
      const statusStore = useStatusStore();
      statusStore.index = statuses as any;
      return { args };
    },
    template: `<BoardForm v-bind="args" />`,
  }),
  args: {
    submitButtonText: 'global.create',
    onSubmit: (item: any) => console.log('submit', item),
  },
};

export const Edit: Story = {
  render: (args) => ({
    components: { BoardForm },
    setup() {
      const statusStore = useStatusStore();
      statusStore.index = statuses as any;
      return { args };
    },
    template: `<BoardForm v-bind="args" />`,
  }),
  args: {
    submitButtonText: 'global.save',
    onSubmit: (item: any) => console.log('submit', item),
    onDelete: () => console.log('delete'),
    initialData: { id: '1', name: 'Dev Board', description: 'Main development board', boardColumns: [] },
  },
};
