<template>
  <priority-table
    :items="priorityStore.index"
    :onEdit="handleEdit"
  />
</template>

<script lang="ts" setup>
import { usePriorityStore } from '@/store/priority.store';
import PriorityTable from '@/component/PriorityTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';
import { useRouter } from 'vue-router';
import { Priority } from '@/lib/interfaces/priority.interface';

const router = useRouter();
const layoutStore = useLayoutStore();
const priorityStore = usePriorityStore();

const handleEdit = (priority: Priority) => {
  router.push(`/admin/priority/${priority.id}`);
};

onMounted(async () => {
  await priorityStore.getPriorities();
  layoutStore.setActions([
    {
      text: 'global.create',
      to: '/admin/priority/create',
      color: 'primary',
      icon: 'mdi-plus',
    },
  ]);
});

onUnmounted(() => {
  layoutStore.resetActions();
});
</script>
