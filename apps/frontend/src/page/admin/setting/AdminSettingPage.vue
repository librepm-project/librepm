<template>
  <v-container fluid>
    <v-form
      ref="form"
      @submit.prevent="saveSettings"
    >
      <v-row v-if="settingStore.loading || loadingData">
        <v-col
          cols="12"
          class="text-center"
        >
          <v-progress-circular
            indeterminate
            color="primary"
          />
        </v-col>
      </v-row>
      <v-row v-else-if="settingStore.error">
        <v-col
          cols="12"
          class="text-center text-red"
        >
          {{ settingStore.error }}
        </v-col>
      </v-row>
      <template v-else-if="settingStore.settings.filter(s => s.key && s.key !== 'undefined').length > 0">
        <v-row
          v-for="setting in settingStore.settings.filter(s => s.key && s.key !== 'undefined')"
          :key="setting.key"
        >
          <v-col cols="12">
            <!-- Custom dropdowns for special default settings -->
            <v-select
              v-if="setting.key === 'default_board_id'"
              v-model="setting.value"
              :label="$t('settings.' + setting.key)"
              :items="boards"
              item-title="name"
              item-value="id"
              variant="outlined"
              clearable
            />
            <v-select
              v-else-if="setting.key === 'default_dashboard_id'"
              v-model="setting.value"
              :label="$t('settings.' + setting.key)"
              :items="dashboards"
              item-title="name"
              item-value="id"
              variant="outlined"
              clearable
            />

            <!-- Standard inputs for other settings -->
            <v-switch
              v-else-if="setting.valueType === 'boolean'"
              v-model="setting.value"
              :label="$t('settings.' + setting.key)"
              color="primary"
              hide-details
            />
            <v-select
              v-else-if="setting.options && setting.options.length > 0"
              v-model="setting.value"
              :label="$t('settings.' + setting.key)"
              :items="setting.options"
              variant="outlined"
              :rules="[requiredRule]"
            />
            <v-text-field
              v-else
              v-model="setting.value"
              :label="$t('settings.' + setting.key)"
              variant="outlined"
              :rules="setting.valueType === 'json' ? [] : [requiredRule]"
            />
          </v-col>
        </v-row>
        <v-row class="mt-2">
          <v-col cols="12">
            <v-btn
              type="submit"
              color="primary"
              size="large"
              prepend-icon="mdi-content-save"
              rounded="lg"
              class="font-weight-bold"
              :loading="settingStore.loading"
            >
              {{ $t('global.save') }}
            </v-btn>
          </v-col>
        </v-row>
      </template>
      <v-row v-else>
        <v-col
          cols="12"
          class="text-center"
        >
          {{ $t('settings.no_settings_found') }}
        </v-col>
      </v-row>
    </v-form>
  </v-container>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useSettingStore } from '@/store/setting.store';
import { requiredRule } from '@/lib/formRule';
import boardApi from '@/api/board.api';
import dashboardApi from '@/api/dashboard.api';
import { Board } from '@/lib/interfaces/board.interface';
import { Dashboard } from '@/lib/interfaces/dashboard.interface';

const settingStore = useSettingStore();
const form = ref(null);
const loadingData = ref(false);

const boards = ref<Board[]>([]);
const dashboards = ref<Dashboard[]>([]);

const saveSettings = async () => {
  const { valid } = await (form.value as any).validate();
  if (valid) {
    const validSettings = settingStore.settings.filter(s => s.key && s.key !== 'undefined');
    for (const setting of validSettings) {
      let valueToSend = setting.value || "";
      if (setting.valueType === 'boolean') {
        valueToSend = String(setting.value);
      }
      await settingStore.update(setting.key, valueToSend);
    }
    await settingStore.fetch();
  }
};

onMounted(async () => {
  loadingData.value = true;
  try {
    const [fetchedBoards, fetchedDashboards] = await Promise.all([
      boardApi.index(),
      dashboardApi.index(),
    ]);
    await settingStore.fetch();
    boards.value = fetchedBoards;
    dashboards.value = fetchedDashboards;
  } finally {
    loadingData.value = false;
  }
});
</script>
