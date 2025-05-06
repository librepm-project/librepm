<template>
  <div>
    <div class="text-right">
      <v-btn :text="t('global.create')" to="/admin/status/create" color="primary" prepend-icon="mdi-plus"></v-btn>
    </div>
    <status-table :items="statusStore.index" />
  </div>
</template>

<script lang="ts" setup>
  import { useStatusStore } from '@/store/status.store';
  import StatusTable from '@/component/StatusTable.vue';
  import { onMounted } from 'vue';
  import { useI18n } from 'vue-i18n'

  const { t } = useI18n();

  const statusStore = useStatusStore();

  onMounted(async () => {
    await statusStore.getStatuses();
  });

</script>
