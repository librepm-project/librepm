import './styles.scss';
import 'vuetify/styles'

import { createApp } from 'vue';
import { createVuetify } from 'vuetify'
import { createPinia } from 'pinia'

import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

import App from './app/App.vue';
import router from './router/index.router';

const vuetify = createVuetify({
  components,
  directives,
})

const pinia = createPinia()

const app = createApp(App);
app.use(pinia);
app.use(vuetify);
app.use(router);
app.mount('#root');
