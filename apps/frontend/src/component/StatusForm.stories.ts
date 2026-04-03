import type { Meta, StoryObj } from '@storybook/vue3';
import StatusForm from './StatusForm.vue';

const meta: Meta<typeof StatusForm> = {
  component: StatusForm,
  title: 'Components/StatusForm',
};
export default meta;

type Story = StoryObj<typeof StatusForm>;

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
    status: { id: '1', name: 'In Progress', color: '#FFC107' },
  },
};
