import 'vuetify/styles';
import '@/theme/styles.scss';

import { createApp } from 'vue';
import { createVuetify } from 'vuetify';
import { createPinia } from 'pinia';

import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';

import App from '@/layout/App.vue';
import router from '@/router/index.router';
import { theme } from '@/theme/theme';
import { i18n } from '@/i18n';
import '@mdi/font/css/materialdesignicons.css'
import { useAppConfigStore } from '@/store/app-config.store';
import { useSettingStore } from '@/store/setting.store';

const vuetify = createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: 'librepm',
    themes: {
      librepm: theme,
    },
  },
  defaults: {
    VChip: {
      variant: 'outlined',
      class: 'border-opacity-100',
    },
    VTextField: {
      variant: 'outlined',
      density: 'compact',
    },
    VTextarea: {
      variant: 'outlined',
      density: 'compact',
    },
    VSelect: {
      variant: 'outlined',
      density: 'compact',
    },
    VDataTable: {
      class: 'border-thin rounded-xl',
      density: 'compact',
    },
    VList: {
      density: 'compact',
    },
    VListItem: {
      density: 'compact',
    },
    VBtn: {
      density: 'comfortable',
    },
  },
});

const pinia = createPinia();

const app = createApp(App);
app.use(i18n);
app.use(pinia);
app.use(vuetify);
app.use(router);

const configStore = useAppConfigStore();
const settingStore = useSettingStore();
Promise.allSettled([
  configStore.fetch(),
  settingStore.fetch(),
]).finally(() => {
  app.mount('#root');
});
