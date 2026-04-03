import type { Meta, StoryObj } from '@storybook/vue3';
import ProjectReleaseTable from './ProjectReleaseTable.vue';

const meta: Meta<typeof ProjectReleaseTable> = {
  component: ProjectReleaseTable,
  title: 'Components/ProjectReleaseTable',
};
export default meta;

type Story = StoryObj<typeof ProjectReleaseTable>;

const items = [
  {
    id: '1',
    release_id: 'r1',
    project_id: 'p1',
    release: { id: 'r1', name: 'v1.0.0', status: 'released', released_at: '2025-01-15' },
    project: { id: 'p1', name: 'Wedding Project', codeName: 'WED' },
  },
  {
    id: '2',
    release_id: 'r1',
    project_id: 'p2',
    release: { id: 'r1', name: 'v1.0.0', status: 'released', released_at: '2025-01-15' },
    project: { id: 'p2', name: 'Startup Launch', codeName: 'START' },
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
