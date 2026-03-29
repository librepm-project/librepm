<template>
  <div class="mb-6">
    <input
      ref="fileInputRef"
      type="file"
      hidden
      @change="onFileSelected"
    >

    <!-- Fejléc -->
    <div class="d-flex align-center mb-3">
      <p class="text-subtitle2 font-weight-medium mb-0 flex-grow-1">
        <v-icon
          x-small
          class="mr-1"
        >
          mdi-paperclip
        </v-icon>
        {{ t('attachment.attachments') }}
      </p>
      <v-btn
        icon
        size="x-small"
        variant="text"
        @click="triggerFileInput"
      >
        <v-icon>mdi-plus</v-icon>
      </v-btn>
    </div>

    <!-- Lista -->
    <div
      v-for="attachment in attachments"
      :key="attachment.id"
      class="d-flex align-center gap-2 pa-2 rounded hover-bg mb-2"
    >
      <v-icon
        size="small"
        class="text-disabled"
      >
        mdi-file-outline
      </v-icon>
      <div class="flex-grow-1 min-width-0">
        <span class="text-body2">{{ attachment.filename }}</span>
        <span class="text-caption text-disabled ml-2">{{ formatSize(attachment.size) }}</span>
      </div>
      <v-btn
        icon
        size="x-small"
        variant="text"
        @click="downloadAttachment(attachment)"
      >
        <v-icon size="small">
          mdi-download
        </v-icon>
      </v-btn>
      <v-btn
        icon
        size="x-small"
        variant="text"
        color="error"
        @click="remove(attachment.id!)"
      >
        <v-icon size="small">
          mdi-close
        </v-icon>
      </v-btn>
    </div>

    <p
      v-if="attachments.length === 0"
      class="text-caption text-disabled mb-0"
    >
      {{ t('attachment.no_attachments') }}
    </p>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { Attachment } from '@/lib/interfaces/attachment.interface';
import { useAttachmentStore } from '@/store/attachment.store';
import attachmentApi from '@/api/attachment.api';

interface Props {
  issueId: string;
  attachments: Attachment[];
}

const props = defineProps<Props>();
const { t } = useI18n();
const attachmentStore = useAttachmentStore();
const fileInputRef = ref<HTMLInputElement | null>(null);

function formatSize(bytes: number): string {
  if (bytes >= 1024 * 1024) return (bytes / (1024 * 1024)).toFixed(1) + ' MB';
  if (bytes >= 1024) return (bytes / 1024).toFixed(1) + ' KB';
  return bytes + ' B';
}

const triggerFileInput = () => {
  fileInputRef.value?.click();
};

const onFileSelected = async (event: Event) => {
  const input = event.target as HTMLInputElement;
  if (!input.files?.length) return;
  const file = input.files[0];
  await attachmentStore.create(props.issueId, file);
  input.value = '';
};

const downloadAttachment = async (attachment: Attachment) => {
  await attachmentApi.download(attachment.id!, attachment.filename);
};

const remove = async (attachmentId: string) => {
  await attachmentStore.destroy(attachmentId);
};
</script>

<style scoped>
.gap-2 {
  gap: 0.5rem;
}

.hover-bg:hover {
  background-color: rgba(63, 81, 181, 0.05);
}

.min-width-0 {
  min-width: 0;
}
</style>
