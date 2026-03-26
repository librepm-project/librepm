<template>
  <v-container fluid>
    <v-card elevation="2" class="rounded-xl overflow-hidden">
      <v-card-title class="bg-primary text-white d-flex justify-space-between align-center py-2 px-4">
        <div class="d-flex align-center gap-2 flex-grow-1 mr-4">
          <v-icon size="small">mdi-cog</v-icon>
          <div class="text-h6 m-0 flex-grow-1 font-weight-bold">{{ $t('admin.settings') }}</div>
        </div>
      </v-card-title>

      <v-divider />

      <v-card-text class="pa-4">
        <v-form @submit.prevent="saveSettings" ref="form">
          <v-row v-if="settingStore.loading">
            <v-col cols="12" class="text-center">
              <v-progress-circular indeterminate color="primary"></v-progress-circular>
            </v-col>
          </v-row>
          <v-row v-else-if="settingStore.error">
            <v-col cols="12" class="text-center text-red">{{ settingStore.error }}</v-col>
          </v-row>
          <!-- Filter settings to only include those with a valid key -->
          <v-row v-else-if="settingStore.settings.filter(s => s.key && s.key !== 'undefined').length > 0">
            <v-col
              v-for="setting in settingStore.settings.filter(s => s.key && s.key !== 'undefined')"
              :key="setting.key"
              cols="12" md="6"
            >
              <!-- Render a switch only if valueType is explicitly boolean -->
              <v-switch
                v-if="setting.valueType === 'boolean'"
                v-model="setting.value"
                :label="$t('settings.' + setting.key)"
                color="primary"
                hide-details
                class="mt-4"
              />
              <!-- Otherwise, render a text field, including for strings and cases where valueType is missing -->
              <v-text-field
                v-else
                v-model="setting.value"
                :label="$t('settings.' + setting.key)"
                variant="outlined"
                dense
                class="rounded-lg"
                :rules="[requiredRule]"
              />
            </v-col>
          </v-row>
          <v-row v-else>
            <v-col cols="12" class="text-center">{{ $t('settings.no_settings_found') }}</v-col>
          </v-row>

          <v-row v-if="settingStore.settings.length > 0" class="mt-6">
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
        </v-form>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useSettingStore } from '@/store/setting.store';
// Assuming requiredRule is correctly imported and defined in '@/lib/formRule'
import { requiredRule } from '@/lib/formRule'; 

const settingStore = useSettingStore();
const form = ref(null);

const saveSettings = async () => {
  const { valid } = await (form.value as any).validate();
  if (valid) {
    // Filter settings again here to ensure we only try to save valid ones
    const validSettingsToSave = settingStore.settings.filter(s => s.key && s.key !== 'undefined');
    for (const setting of validSettingsToSave) {
      let valueToSend = setting.value;
      if (setting.valueType === 'boolean') {
        valueToSend = String(setting.value); // Convert boolean to string "true" or "false"
      }
      await settingStore.updateSetting(setting.key, valueToSend);
    }
    await settingStore.fetchSettings();
  }
};

onMounted(() => {
  settingStore.fetchSettings();
});

</script>

<style scoped>
.gap-2 {
  gap: 0.5rem;
}

/* Add styles for v-switch if needed */
.v-switch {
  margin-top: 0 !important; /* Adjust as needed */
  padding-top: 0 !important;
}
</style>
