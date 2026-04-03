import type { Meta, StoryObj } from '@storybook/vue3';
import ProjectTable from './ProjectTable.vue';

const meta: Meta<typeof ProjectTable> = {
  component: ProjectTable,
  title: 'Components/ProjectTable',
};
export default meta;

type Story = StoryObj<typeof ProjectTable>;

const items = [
  { id: '1', name: 'Wedding Project', codeName: 'WED', description: 'Planning the big day', defaultStatusId: '1', defaultTrackerId: '1' },
  { id: '2', name: 'Startup Launch', codeName: 'START', description: 'Founding a new company', defaultStatusId: null, defaultTrackerId: null },
  { id: '3', name: 'London Trip', codeName: 'LON', description: 'Sightseeing in London', defaultStatusId: '2', defaultTrackerId: '2' },
];

export const Default: Story = {
  args: {
    items,
    onEdit: (item: any) => console.log('edit', item),
  },
};

export const Empty: Story = {
  args: {
    items: [],
    onEdit: () => {},
  },
};
