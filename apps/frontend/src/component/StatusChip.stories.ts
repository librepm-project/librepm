import type { Meta, StoryObj } from '@storybook/vue3';
import StatusChip from './StatusChip.vue';

const meta: Meta<typeof StatusChip> = {
  component: StatusChip,
  title: 'Components/StatusChip',
  argTypes: {
    status: { control: 'object' },
  },
};
export default meta;

type Story = StoryObj<typeof StatusChip>;

export const Open: Story = {
  args: {
    status: { id: '1', name: 'Open', color: '#2196F3' },
  },
};

export const InProgress: Story = {
  args: {
    status: { id: '2', name: 'In Progress', color: '#FFC107' },
  },
};

export const Done: Story = {
  args: {
    status: { id: '3', name: 'Done', color: '#4CAF50' },
  },
};

export const Blocked: Story = {
  args: {
    status: { id: '4', name: 'Blocked', color: '#F44336' },
  },
};

export const Planned: Story = {
  args: {
    status: { id: '5', name: 'Planned', color: '#9E9E9E' },
  },
};
