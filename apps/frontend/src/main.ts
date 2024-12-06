import '@/theme/styles.scss';
import 'vuetify/styles'

import { createApp } from 'vue';
import { createVuetify } from 'vuetify'
import { createPinia } from 'pinia'

import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

import App from '@/layout/App.vue';
import router from '@/router/index.router';
import { themes } from '@/theme/theme';
import { i18n } from '@/i18n';

const vuetify = createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: 'dark',
    themes,
  },
})

const pinia = createPinia()

const app = createApp(App);
app.use(i18n);
app.use(pinia);
app.use(vuetify);
app.use(router);
app.mount('#root');
