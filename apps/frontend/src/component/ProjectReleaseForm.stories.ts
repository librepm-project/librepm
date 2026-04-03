import type { Meta, StoryObj } from '@storybook/vue3';
import ProjectReleaseForm from './ProjectReleaseForm.vue';

const meta: Meta<typeof ProjectReleaseForm> = {
  component: ProjectReleaseForm,
  title: 'Components/ProjectReleaseForm',
};
export default meta;

type Story = StoryObj<typeof ProjectReleaseForm>;

const releases = [
  { id: 'r1', name: 'v1.0.0', status: 'released', released_at: '2025-01-15' },
  { id: 'r2', name: 'v1.1.0', status: 'planned', released_at: null },
];

const projects = [
  { id: 'p1', name: 'Wedding Project', codeName: 'WED' },
  { id: 'p2', name: 'Startup Launch', codeName: 'START' },
];

export const Default: Story = {
  args: {
    submitButtonText: 'global.create',
    onSubmit: async (item: any) => console.log('submit', item),
    releases,
    projects,
  },
};

export const WithPreselectedRelease: Story = {
  args: {
    submitButtonText: 'global.create',
    onSubmit: async (item: any) => console.log('submit', item),
    releaseId: 'r1',
    releases,
    projects,
  },
};
