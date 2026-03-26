<template>
  <div class="pa-1">

    <!-- Lista -->
    <div
      v-for="comment in comments"
      :key="comment.id"
      class="mb-4"
    >
      <!-- Szerkesztés mód -->
      <div v-if="editingId === comment.id" class="pa-3 rounded edit-bg">
        <v-textarea
          v-model="editContent"
          density="compact"
          variant="underlined"
          hide-details
          auto-grow
          rows="2"
          class="mb-2"
        />
        <div class="d-flex gap-2">
          <v-btn size="small" color="primary" :disabled="!editContent.trim()" @click="submitEdit(comment.id!)">
            {{ t('comment.save') }}
          </v-btn>
          <v-btn size="small" variant="text" @click="editingId = null">
            {{ t('global.cancel') }}
          </v-btn>
        </div>
      </div>

      <!-- Olvasás mód -->
      <div v-else class="d-flex align-start gap-2 pa-2 rounded hover-bg">
        <div class="flex-grow-1">
          <div class="d-flex align-center gap-2 mb-1">
            <span class="text-body2 font-weight-medium">{{ userName(comment) }}</span>
            <span class="text-caption text-disabled">{{ formatDate(comment.createdAt) }}</span>
            <span v-if="comment.updatedAt && comment.updatedAt !== comment.createdAt" class="text-caption text-disabled">
              (edited)
            </span>
          </div>
          <p class="text-body2 mb-0" style="white-space: pre-wrap">{{ comment.content }}</p>
        </div>
        <v-btn icon size="x-small" variant="text" @click="startEdit(comment)">
          <v-icon size="small">mdi-pencil</v-icon>
        </v-btn>
        <v-btn icon size="x-small" variant="text" color="error" @click="remove(comment.id!)">
          <v-icon size="small">mdi-close</v-icon>
        </v-btn>
      </div>
    </div>

    <p v-if="comments.length === 0" class="text-caption text-disabled mb-4">
      {{ t('comment.no_comments') }}
    </p>

    <!-- Hozzáadás form -->
    <div class="mt-2">
      <v-textarea
        v-model="newContent"
        :placeholder="t('comment.add_comment')"
        density="compact"
        variant="outlined"
        hide-details
        auto-grow
        rows="2"
        class="mb-2"
      />
      <v-btn
        size="small"
        color="primary"
        :disabled="!newContent.trim()"
        @click="submitNew"
      >
        {{ t('comment.add_comment') }}
      </v-btn>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { Comment } from '@/lib/interfaces/comment.interface';
import { useCommentStore } from '@/store/comment.store';

interface Props {
  issueId: string;
  comments: Comment[];
}

const props = defineProps<Props>();
const { t } = useI18n();
const commentStore = useCommentStore();

const newContent = ref('');
const editingId = ref<string | null>(null);
const editContent = ref('');

function userName(comment: Comment): string {
  if (!comment.user) return 'Unknown';
  return `${comment.user.firstName} ${comment.user.lastName}`.trim() || comment.user.email;
}

function formatDate(iso?: string): string {
  if (!iso) return '';
  return new Date(iso).toLocaleString();
}

const submitNew = async () => {
  if (!newContent.value.trim()) return;
  await commentStore.create(props.issueId, newContent.value.trim());
  newContent.value = '';
};

const startEdit = (comment: Comment) => {
  editingId.value = comment.id!;
  editContent.value = comment.content;
};

const submitEdit = async (commentId: string) => {
  if (!editContent.value.trim()) return;
  await commentStore.update(commentId, editContent.value.trim());
  editingId.value = null;
};

const remove = async (commentId: string) => {
  await commentStore.destroy(commentId);
};
</script>

<style scoped>
.gap-2 {
  gap: 0.5rem;
}

.hover-bg:hover {
  background-color: rgba(63, 81, 181, 0.05);
}

.edit-bg {
  background-color: rgba(0, 0, 0, 0.03);
}
</style>
