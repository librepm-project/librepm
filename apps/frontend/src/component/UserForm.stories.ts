import type { Meta, StoryObj } from '@storybook/vue3';
import UserForm from './UserForm.vue';

const meta: Meta<typeof UserForm> = {
  component: UserForm,
  title: 'Components/UserForm',
};
export default meta;

type Story = StoryObj<typeof UserForm>;

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
    user: {
      id: '1',
      email: 'john@example.com',
      firstName: 'John',
      lastName: 'Doe',
      phone: '123-456-7890',
      language: 'en',
      country: 'US',
      blocked: false,
    },
  },
};
