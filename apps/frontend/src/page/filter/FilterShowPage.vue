<template>
  <div>
    <filter-form
      :filter="filterStore.current"
      :onSubmit="handleSubmit"
      :onDelete="handleDelete"
      submitButtonText="global.update"
    />
  </div>
</template>

<script lang="ts" setup>
import { useFilterStore } from '@/store/filter.store';
import { useRoute, useRouter } from 'vue-router';
import FilterForm from '@/component/FilterForm.vue';
import { Filter } from '@/lib/interfaces/filter.interface';
import { onMounted } from 'vue';

const route = useRoute();
const router = useRouter();
const filterStore = useFilterStore();

const handleSubmit = async (filter: Omit<Filter, 'id'>) => {
  try {
    const filterId = route.params.filterId as string;
    await filterStore.put(filterId, filter);
    router.push(`/filter/${filterId}`);
  } catch (error) {
    console.error('Error updating filter:', error);
  }
};

const handleDelete = async () => {
  if (confirm('Are you sure you want to delete this filter?')) {
    try {
      const filterId = route.params.filterId as string;
      await filterStore.delete(filterId);
      router.push('/filter');
    } catch (error) {
      console.error('Error deleting filter:', error);
    }
  }
};

onMounted(async () => {
  filterStore.current = null;
  const filterId = route.params.filterId as string;
  await filterStore.get(filterId);
});
</script>
