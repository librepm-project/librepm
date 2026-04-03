import type { Meta, StoryObj } from '@storybook/vue3';
import InlineSidebarEdit from './InlineSidebarEdit.vue';
import StatusChip from './StatusChip.vue';

const meta: Meta<typeof InlineSidebarEdit> = {
  component: InlineSidebarEdit,
  title: 'Components/InlineSidebarEdit',
  argTypes: {
    label: { control: 'text' },
    icon: { control: 'text' },
    modelValue: { control: 'text' },
    clearable: { control: 'boolean' },
  },
};
export default meta;

type Story = StoryObj<typeof InlineSidebarEdit>;

const statusItems = [
  { id: '1', name: 'Open', color: '#2196F3' },
  { id: '2', name: 'In Progress', color: '#FFC107' },
  { id: '3', name: 'Done', color: '#4CAF50' },
];

export const WithValue: Story = {
  args: {
    label: 'Status',
    icon: 'mdi-circle-slice-8',
    modelValue: '1',
    items: statusItems,
    itemTitle: 'name',
    itemValue: 'id',
    clearable: false,
  },
  render: (args) => ({
    components: { InlineSidebarEdit, StatusChip },
    setup: () => ({ args, statusItems }),
    template: `
      <div style="width: 240px; padding: 16px;">
        <InlineSidebarEdit v-bind="args">
          <StatusChip :status="statusItems.find(s => s.id === args.modelValue) ?? statusItems[0]" />
        </InlineSidebarEdit>
      </div>
    `,
  }),
};

export const Clearable: Story = {
  args: {
    label: 'Assignee',
    icon: 'mdi-account',
    modelValue: null,
    items: [
      { id: '1', name: 'John Doe' },
      { id: '2', name: 'Jane Doe' },
    ],
    itemTitle: 'name',
    itemValue: 'id',
    clearable: true,
  },
  render: (args) => ({
    components: { InlineSidebarEdit },
    setup: () => ({ args }),
    template: `
      <div style="width: 240px; padding: 16px;">
        <InlineSidebarEdit v-bind="args">
          <span class="text-body-2 text-medium-emphasis">— unassigned —</span>
        </InlineSidebarEdit>
      </div>
    `,
  }),
};
