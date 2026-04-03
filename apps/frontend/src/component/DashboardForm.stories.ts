import type { Meta, StoryObj } from '@storybook/vue3';
import DashboardForm from './DashboardForm.vue';

const meta: Meta<typeof DashboardForm> = {
  component: DashboardForm,
  title: 'Components/DashboardForm',
};
export default meta;

type Story = StoryObj<typeof DashboardForm>;

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
  },
};
