import { defineStore } from 'pinia';
import { Setting } from '@/lib/interfaces/setting.interface';
import settingApi from '@/api/setting.api';

interface SettingState {
    settings: Setting[];
    loading: boolean;
    error: string | null;
}

export const useSettingStore = defineStore('setting', {
    state: (): SettingState => ({
        settings: [],
        loading: false,
        error: null,
    }),

    actions: {
        async fetchSettings() {
            this.loading = true;
            this.error = null;
            try {
                const data = await settingApi.index();
                this.settings = data.map((item: any) => ({
                    key: item.Key || item.key,
                    value: item.Value || item.value,
                    valueType: item.ValueType || item.valueType,
                    options: item.Options || item.options,
                }));
            } catch (error: any) {
                this.error = error.message || 'Failed to fetch settings';
                console.error('Error fetching settings:', error);
            } finally {
                this.loading = false;
            }
        },

        async updateSetting(key: string, value: string) {
            this.loading = true;
            this.error = null;
            try {
                await settingApi.update(key, value);
                // Optimistic update in state
                const setting = this.settings.find(s => s.key === key);
                if (setting) {
                    setting.value = value;
                }
            } catch (error: any) {
                this.error = error.message || 'Failed to update setting';
                console.error('Error updating setting:', error);
            } finally {
                this.loading = false;
            }
        },

        // Helper to get a setting value by key
        getSettingValue(key: string): string | undefined {
            const setting = this.settings.find(s => s.key === key);
            return setting?.value;
        },
    },
});
