import type { Meta, StoryObj } from '@storybook/vue3';
import ProjectReleaseIssueTable from './ProjectReleaseIssueTable.vue';

const meta: Meta<typeof ProjectReleaseIssueTable> = {
  component: ProjectReleaseIssueTable,
  title: 'Components/ProjectReleaseIssueTable',
};
export default meta;

type Story = StoryObj<typeof ProjectReleaseIssueTable>;

const items = [
  {
    id: '1',
    project_release_id: 'pr1',
    issue_id: 'i1',
    issue: {
      id: 'i1',
      key: 'WED-1',
      summary: 'Book Venue',
      description: 'Find a venue',
      status: { id: '1', name: 'Done', color: '#4CAF50' },
      tracker: { id: '1', name: 'Task', color: '#6C63FF' },
    },
  },
  {
    id: '2',
    project_release_id: 'pr1',
    issue_id: 'i2',
    issue: {
      id: 'i2',
      key: 'WED-2',
      summary: 'Choose Menu',
      description: 'Select the dinner menu',
      status: { id: '2', name: 'In Progress', color: '#FFC107' },
      tracker: { id: '1', name: 'Task', color: '#6C63FF' },
    },
  },
];

export const Default: Story = {
  args: {
    items,
    onDelete: (id: any) => console.log('delete', id),
  },
};

export const Empty: Story = {
  args: {
    items: [],
    onDelete: () => {},
  },
};
