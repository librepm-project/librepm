import type { Meta, StoryObj } from '@storybook/vue3';
import PriorityChip from './PriorityChip.vue';

const meta: Meta<typeof PriorityChip> = {
  component: PriorityChip,
  title: 'Components/PriorityChip',
  argTypes: {
    priority: { control: 'object' },
  },
};
export default meta;

type Story = StoryObj<typeof PriorityChip>;

export const High: Story = {
  args: {
    priority: { id: '1', name: 'High', color: '#F44336' },
  },
};

export const Normal: Story = {
  args: {
    priority: { id: '2', name: 'Normal', color: '#FFC107' },
  },
};

export const Low: Story = {
  args: {
    priority: { id: '3', name: 'Low', color: '#4CAF50' },
  },
};
