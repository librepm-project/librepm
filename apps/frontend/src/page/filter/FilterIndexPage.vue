<template>
  <div>
    <div class="text-right">
      <v-btn :text="t('global.create')" to="/filter/create" color="primary" prepend-icon="mdi-plus"></v-btn>
    </div>
    <filter-table :items="filterStore.index" />
  </div>
</template>

<script lang="ts" setup>
  import { useFilterStore } from '@/store/filter.store';
  import FilterTable from '@/component/FilterTable.vue';
  import { onMounted } from 'vue';
  import { useI18n } from 'vue-i18n'

const { t } = useI18n();

  const filterStore = useFilterStore();

  onMounted(async () => {
    await filterStore.getFilters()
  })

</script>