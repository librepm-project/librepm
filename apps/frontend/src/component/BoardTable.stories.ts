import type { Meta, StoryObj } from '@storybook/vue3';
import BoardTable from './BoardTable.vue';

const meta: Meta<typeof BoardTable> = {
  component: BoardTable,
  title: 'Components/BoardTable',
};
export default meta;

type Story = StoryObj<typeof BoardTable>;

const items = [
  { id: '1', name: 'Dev Board', description: 'Main development board', boardColumns: [] },
  { id: '2', name: 'Wedding Board', description: 'Wedding planning board', boardColumns: [] },
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
