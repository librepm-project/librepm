import type { Meta, StoryObj } from '@storybook/vue3';
import TrackerForm from './TrackerForm.vue';

const meta: Meta<typeof TrackerForm> = {
  component: TrackerForm,
  title: 'Components/TrackerForm',
};
export default meta;

type Story = StoryObj<typeof TrackerForm>;

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
    tracker: { id: '1', name: 'Bug', color: '#FF4B4B' },
  },
};
