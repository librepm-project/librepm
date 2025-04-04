<template>
  <div>
    <dashboard-form :onSubmit="dashboardStore.postDashboard" submitButtonText="global.create" />
  </div>
</template>

<script lang="ts" setup>
import { useDashboardStore } from '@/store/dashboard.store';

import DashboardForm from '@/component/DashboardForm.vue';

const dashboardStore = useDashboardStore();

</script>