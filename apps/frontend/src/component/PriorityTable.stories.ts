import type { Meta, StoryObj } from '@storybook/vue3';
import PriorityTable from './PriorityTable.vue';

const meta: Meta<typeof PriorityTable> = {
  component: PriorityTable,
  title: 'Components/PriorityTable',
};
export default meta;

type Story = StoryObj<typeof PriorityTable>;

const items = [
  { id: '1', name: 'High', color: '#F44336' },
  { id: '2', name: 'Normal', color: '#FFC107' },
  { id: '3', name: 'Low', color: '#4CAF50' },
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
