import type { Meta, StoryObj } from '@storybook/vue3';
import AttachmentPanel from './AttachmentPanel.vue';

const meta: Meta<typeof AttachmentPanel> = {
  component: AttachmentPanel,
  title: 'Components/AttachmentPanel',
};
export default meta;

type Story = StoryObj<typeof AttachmentPanel>;

const attachments = [
  {
    id: 'a1',
    filename: 'venue-brochure.pdf',
    mimeType: 'application/pdf',
    size: 2048576,
    entityType: 'issue',
    entityId: 'i1',
    createdAt: '2025-03-01T10:00:00Z',
  },
  {
    id: 'a2',
    filename: 'floor-plan.png',
    mimeType: 'image/png',
    size: 512000,
    entityType: 'issue',
    entityId: 'i1',
    createdAt: '2025-03-02T14:30:00Z',
  },
];

export const Default: Story = {
  args: {
    issueId: 'i1',
    attachments,
  },
};

export const Empty: Story = {
  args: {
    issueId: 'i1',
    attachments: [],
  },
};
