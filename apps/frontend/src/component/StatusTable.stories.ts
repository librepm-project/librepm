import type { Meta, StoryObj } from '@storybook/vue3';
import StatusTable from './StatusTable.vue';

const meta: Meta<typeof StatusTable> = {
  component: StatusTable,
  title: 'Components/StatusTable',
};
export default meta;

type Story = StoryObj<typeof StatusTable>;

const items = [
  { id: '1', name: 'Open', color: '#2196F3' },
  { id: '2', name: 'In Progress', color: '#FFC107' },
  { id: '3', name: 'Done', color: '#4CAF50' },
  { id: '4', name: 'Blocked', color: '#F44336' },
];

export const Default: Story = {
  args: {
    items,
    onEdit: (item: any) => console.log('edit', item),
  },
};

export const Empty: Story = {
  args: {
    items: [],
    onEdit: () => {},
  },
};
