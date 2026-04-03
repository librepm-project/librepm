import type { Meta, StoryObj } from '@storybook/vue3';
import FilterForm from './FilterForm.vue';

const meta: Meta<typeof FilterForm> = {
  component: FilterForm,
  title: 'Components/FilterForm',
};
export default meta;

type Story = StoryObj<typeof FilterForm>;

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
    filter: {
      id: '1',
      name: 'Open Issues',
      description: 'All issues that are not closed',
      conditions: [],
    },
  },
};
