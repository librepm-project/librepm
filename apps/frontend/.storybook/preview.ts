import { setup, type Preview } from '@storybook/vue3';
import { createVuetify } from 'vuetify';
import { createPinia } from 'pinia';
import { createRouter, createWebHashHistory } from 'vue-router';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';
import 'vuetify/styles';
import '@mdi/font/css/materialdesignicons.css';
import '../src/theme/styles.scss';
import { theme } from '../src/theme/theme';
import { i18n } from '../src/i18n';

setup((app) => {
  const vuetify = createVuetify({
    components,
    directives,
    theme: {
      defaultTheme: 'librepm',
      themes: { librepm: theme },
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

  const router = createRouter({
    history: createWebHashHistory(),
    routes: [{ path: '/:pathMatch(.*)*', component: { template: '<div />' } }],
  });

  app.use(vuetify);
  app.use(createPinia());
  app.use(i18n);
  app.use(router);
});

const preview: Preview = {
  parameters: {
    backgrounds: {
      default: 'app',
      values: [
        { name: 'app', value: '#F5F5F5' },
        { name: 'white', value: '#FFFFFF' },
      ],
    },
  },
};

export default preview;
