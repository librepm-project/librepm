import type { Meta, StoryObj } from '@storybook/vue3';
import RichTextField from './RichTextField.vue';
import { ref } from 'vue';

const meta: Meta<typeof RichTextField> = {
  component: RichTextField,
  title: 'Components/RichTextField',
  argTypes: {
    label: { control: 'text' },
    modelValue: { control: 'text' },
  },
};
export default meta;

type Story = StoryObj<typeof RichTextField>;

export const Empty: Story = {
  args: {
    label: 'Description',
    modelValue: '',
  },
  render: (args) => ({
    components: { RichTextField },
    setup() {
      const content = ref(args.modelValue);
      return { args, content };
    },
    template: `
      <div style="padding: 16px; max-width: 720px;">
        <RichTextField :label="args.label" v-model="content" />
      </div>
    `,
  }),
};

export const WithContent: Story = {
  args: {
    label: 'Description',
    modelValue: '<p>This is a <strong>rich text</strong> description with some <em>formatted</em> content.</p>',
  },
  render: (args) => ({
    components: { RichTextField },
    setup() {
      const content = ref(args.modelValue);
      return { args, content };
    },
    template: `
      <div style="padding: 16px; max-width: 720px;">
        <RichTextField :label="args.label" v-model="content" />
      </div>
    `,
  }),
};
