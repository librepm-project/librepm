import type { Meta, StoryObj } from '@storybook/vue3';
import DashboardFilterWidget from './DashboardFilterWidget.vue';

const meta: Meta<typeof DashboardFilterWidget> = {
  component: DashboardFilterWidget,
  title: 'Components/DashboardFilterWidget',
};
export default meta;

type Story = StoryObj<typeof DashboardFilterWidget>;

export const Default: Story = {
  args: {
    widget: {
      id: 'w1',
      type: 'filter',
      title: 'Open Issues',
      filterId: 'f1',
      filter: { id: 'f1', name: 'Open Issues', description: 'All open issues', conditions: [] },
      weight: 1,
      width: '1/2',
    },
  },
};

export const FullWidth: Story = {
  args: {
    widget: {
      id: 'w2',
      type: 'filter',
      title: 'Wedding Tasks',
      filterId: 'f2',
      filter: { id: 'f2', name: 'Wedding Tasks', description: 'All wedding related tasks', conditions: [] },
      weight: 2,
      width: '1/1',
    },
  },
};
