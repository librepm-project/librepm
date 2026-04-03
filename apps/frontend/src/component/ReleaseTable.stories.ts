import type { Meta, StoryObj } from '@storybook/vue3';
import ReleaseTable from './ReleaseTable.vue';

const meta: Meta<typeof ReleaseTable> = {
  component: ReleaseTable,
  title: 'Components/ReleaseTable',
};
export default meta;

type Story = StoryObj<typeof ReleaseTable>;

const items = [
  { id: '1', name: 'v1.0.0', status: 'released', released_at: '2025-01-15', created_at: '2025-01-01' },
  { id: '2', name: 'v1.1.0', status: 'planned', released_at: null, created_at: '2025-02-01' },
  { id: '3', name: 'v2.0.0', status: 'planned', released_at: null, created_at: '2025-03-01' },
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
