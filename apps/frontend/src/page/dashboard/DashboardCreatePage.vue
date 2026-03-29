<template>
  <div>
    <dashboard-form
      :on-submit="dashboardStore.postDashboard"
      submit-button-text="global.create"
    />
  </div>
</template>

<script lang="ts" setup>
import { useDashboardStore } from '@/store/dashboard.store';

import DashboardForm from '@/component/DashboardForm.vue';

const dashboardStore = useDashboardStore();

</script>