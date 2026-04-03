import type { Meta, StoryObj } from '@storybook/vue3';
import TrackerTable from './TrackerTable.vue';

const meta: Meta<typeof TrackerTable> = {
  component: TrackerTable,
  title: 'Components/TrackerTable',
};
export default meta;

type Story = StoryObj<typeof TrackerTable>;

const items = [
  { id: '1', name: 'Bug', color: '#FF4B4B' },
  { id: '2', name: 'Task', color: '#6C63FF' },
  { id: '3', name: 'Epic', color: '#FFB400' },
  { id: '4', name: 'Story', color: '#00B4B6' },
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
