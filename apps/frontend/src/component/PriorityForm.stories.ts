import type { Meta, StoryObj } from '@storybook/vue3';
import PriorityForm from './PriorityForm.vue';

const meta: Meta<typeof PriorityForm> = {
  component: PriorityForm,
  title: 'Components/PriorityForm',
};
export default meta;

type Story = StoryObj<typeof PriorityForm>;

export const Create: Story = {
  args: {
    submitButtonText: 'global.create',
    onSubmit: (item: any) => console.log('submit', item),
  },
};

export const Edit: Story = {
  args: {
    submitButtonText: 'global.save',
    onSubmit: (item: any) => console.log('submit', item),
    onDelete: () => console.log('delete'),
    priority: { id: '1', name: 'High', color: '#F44336' },
  },
};
