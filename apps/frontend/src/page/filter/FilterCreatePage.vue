<template>
  <div>
    <filter-form :onSubmit="handleSubmit" submitButtonText="global.create" />
  </div>
</template>

<script lang="ts" setup>
import { useFilterStore } from '@/store/filter.store';
import { useRouter } from 'vue-router';
import FilterForm from '@/component/FilterForm.vue';
import { Filter } from '@/lib/interfaces/filter.interface';

const router = useRouter();
const filterStore = useFilterStore();

const handleSubmit = async (filter: Omit<Filter, 'id'>) => {
  try {
    await filterStore.post(filter);
    router.push('/filter');
  } catch (error) {
    console.error('Error creating filter:', error);
  }
};

</script>