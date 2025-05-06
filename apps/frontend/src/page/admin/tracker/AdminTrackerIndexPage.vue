<template>
  <div>
    <div class="text-right">
      <v-btn :text="t('global.create')" to="/admin/tracker/create" color="primary" prepend-icon="mdi-plus"></v-btn>
    </div>
    <tracker-table :items="trackerStore.index" />
  </div>
</template>

<script lang="ts" setup>
  import { useTrackerStore } from '@/store/tracker.store';
  import TrackerTable from '@/component/TrackerTable.vue';
  import { onMounted } from 'vue';
  import { useI18n } from 'vue-i18n'

  const { t } = useI18n();

  const trackerStore = useTrackerStore();

  onMounted(async () => {
    await trackerStore.getTrackers();
  });

</script>
