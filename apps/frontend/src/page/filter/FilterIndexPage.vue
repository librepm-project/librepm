<template>
  <filter-table 
    :items="filterStore.index" 
    :onEdit="handleEdit" 
    :onDelete="handleDelete"
  />
</template>

<script lang="ts" setup>
import { useFilterStore } from '@/store/filter.store';
import FilterTable from '@/component/FilterTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';
import { useRouter } from 'vue-router';
import { Filter } from '@/lib/interfaces/filter.interface';

const router = useRouter();
const layoutStore = useLayoutStore();
const filterStore = useFilterStore();

const handleEdit = (filter: Filter) => {
  router.push(`/filter/${filter.id}`);
};

const handleDelete = async (filter: Filter) => {
  if (confirm('Are you sure?')) {
    try {
      await filterStore.deleteFilter(filter.id!);
    } catch (error) {
      console.error('Error deleting filter:', error);
    }
  }
};

onMounted(async () => {
  await filterStore.getFilters();
  layoutStore.setActions([
    {
      text: 'global.create',
      to: '/filter/create',
      color: 'primary',
      icon: 'mdi-plus'
    },
  ]);
})

onUnmounted(async () => {
  layoutStore.resetActions();
});

</script>