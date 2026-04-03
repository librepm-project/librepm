import type { Meta, StoryObj } from '@storybook/vue3';
import ReleaseForm from './ReleaseForm.vue';

const meta: Meta<typeof ReleaseForm> = {
  component: ReleaseForm,
  title: 'Components/ReleaseForm',
};
export default meta;

type Story = StoryObj<typeof ReleaseForm>;

export const Create: Story = {
  args: {
    submitButtonText: 'global.create',
    onSubmit: async (item: any) => console.log('submit', item),
  },
};

export const Edit: Story = {
  args: {
    submitButtonText: 'global.save',
    onSubmit: async (item: any) => console.log('submit', item),
    initialData: { id: '1', name: 'v1.0.0', status: 'planned', released_at: null },
  },
};
