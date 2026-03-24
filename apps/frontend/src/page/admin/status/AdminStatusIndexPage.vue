<template>
  <status-table 
    :items="statusStore.index" 
    :onEdit="handleEdit" 
    :onDelete="handleDelete"
  />
</template>

<script lang="ts" setup>
import { useStatusStore } from '@/store/status.store';
import StatusTable from '@/component/StatusTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useI18n } from 'vue-i18n'
import { useLayoutStore } from '@/store/layout.store';
import { useRouter } from 'vue-router';
import { Status } from '@/lib/interfaces/status.interface';

const { t } = useI18n();
const router = useRouter();
const layoutStore = useLayoutStore();
const statusStore = useStatusStore();

const handleEdit = (status: Status) => {
  router.push(`/admin/status/${status.id}`);
};

const handleDelete = async (status: Status) => {
  if (confirm('Are you sure?')) {
    try {
      await statusStore.deleteStatus(status.id);
    } catch (error) {
      console.error('Error deleting status:', error);
    }
  }
};

onMounted(async () => {
  await statusStore.getStatuses();
  layoutStore.setActions([
  {
    text: 'global.create',
    to: '/admin/status/create',
    color: 'primary',
    icon: 'mdi-plus'
  }
]);
});

onUnmounted(async () => {
  layoutStore.resetActions();
});
</script>
