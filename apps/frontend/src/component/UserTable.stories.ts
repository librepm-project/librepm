import type { Meta, StoryObj } from '@storybook/vue3';
import UserTable from './UserTable.vue';

const meta: Meta<typeof UserTable> = {
  component: UserTable,
  title: 'Components/UserTable',
};
export default meta;

type Story = StoryObj<typeof UserTable>;

const items = [
  { id: '1', email: 'john@example.com', firstName: 'John', lastName: 'Doe', phone: '123-456-7890', language: 'en', country: 'US', blocked: false },
  { id: '2', email: 'jane@example.com', firstName: 'Jane', lastName: 'Smith', phone: '098-765-4321', language: 'en', country: 'GB', blocked: false },
  { id: '3', email: 'blocked@example.com', firstName: 'Bob', lastName: 'Blocked', phone: '', language: 'en', country: 'US', blocked: true },
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
