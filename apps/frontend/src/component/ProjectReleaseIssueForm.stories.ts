import type { Meta, StoryObj } from '@storybook/vue3';
import ProjectReleaseIssueForm from './ProjectReleaseIssueForm.vue';

const meta: Meta<typeof ProjectReleaseIssueForm> = {
  component: ProjectReleaseIssueForm,
  title: 'Components/ProjectReleaseIssueForm',
};
export default meta;

type Story = StoryObj<typeof ProjectReleaseIssueForm>;

const issues = [
  {
    id: 'i1', key: 'WED-1', summary: 'Book Venue', description: '',
    tracker: { id: '1', name: 'Task', color: '#6C63FF' },
    status: { id: '1', name: 'Done', color: '#4CAF50' },
  },
  {
    id: 'i2', key: 'WED-2', summary: 'Choose Menu', description: '',
    tracker: { id: '1', name: 'Task', color: '#6C63FF' },
    status: { id: '2', name: 'In Progress', color: '#FFC107' },
  },
  {
    id: 'i3', key: 'WED-3', summary: 'Send Invitations', description: '',
    tracker: { id: '2', name: 'Bug', color: '#FF4B4B' },
    status: { id: '3', name: 'Open', color: '#2196F3' },
  },
];

export const Default: Story = {
  args: {
    projectReleaseId: 'pr1',
    issues,
    onSubmit: async (item: any) => console.log('submit', item),
  },
};
