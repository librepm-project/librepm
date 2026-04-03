import type { Meta, StoryObj } from '@storybook/vue3';
import TrackerChip from './TrackerChip.vue';

const meta: Meta<typeof TrackerChip> = {
  component: TrackerChip,
  title: 'Components/TrackerChip',
  argTypes: {
    tracker: { control: 'object' },
  },
};
export default meta;

type Story = StoryObj<typeof TrackerChip>;

export const Bug: Story = {
  args: {
    tracker: { id: '1', name: 'Bug', color: '#FF4B4B' },
  },
};

export const Task: Story = {
  args: {
    tracker: { id: '2', name: 'Task', color: '#6C63FF' },
  },
};

export const Epic: Story = {
  args: {
    tracker: { id: '3', name: 'Epic', color: '#FFB400' },
  },
};

export const Story_: Story = {
  name: 'Story',
  args: {
    tracker: { id: '4', name: 'Story', color: '#00B4B6' },
  },
};
