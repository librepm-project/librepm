import type { Meta, StoryObj } from '@storybook/vue3';
import FilterTable from './FilterTable.vue';

const meta: Meta<typeof FilterTable> = {
  component: FilterTable,
  title: 'Components/FilterTable',
};
export default meta;

type Story = StoryObj<typeof FilterTable>;

const items = [
  { id: '1', name: 'Open Issues', description: 'All issues that are not closed', conditions: [] },
  { id: '2', name: 'My Tasks', description: 'Issues assigned to me', conditions: [] },
  { id: '3', name: 'High Priority', description: 'High priority bugs', conditions: [] },
];

export const Default: Story = {
  args: {
    items,
    onEdit: (item: any) => console.log('edit', item),
    onView: (item: any) => console.log('view', item),
  },
};

export const Empty: Story = {
  args: {
    items: [],
    onEdit: () => {},
    onView: () => {},
  },
};
