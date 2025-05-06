<template>
  <div>
    <status-form :onSubmit="statusStore.postStatus" submitButtonText="global.create" />
  </div>
</template>

<script lang="ts" setup>
import { useStatusStore } from '@/store/status.store';

import StatusForm from '@/component/StatusForm.vue';

const statusStore = useStatusStore();

</script>