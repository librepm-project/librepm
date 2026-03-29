import { defineStore } from 'pinia';
import configApi from '@/api/config.api';

export const useAppConfigStore = defineStore('appConfig', {
    state: () => ({
        registerAllowed: false,
    }),
    actions: {
        async fetch() {
            const config = await configApi.show();
            this.registerAllowed = config.registerAllowed;
        },
    },
});
