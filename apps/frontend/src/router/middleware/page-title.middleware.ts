import {
  NavigationGuardNext,
  RouteLocationNormalized,
  RouteLocationNormalizedLoaded,
} from 'vue-router';

import { reactive } from 'vue';
import { i18n } from '@/i18n';


export const pageTitleMiddleware = (
  to: RouteLocationNormalized,
  _from: RouteLocationNormalizedLoaded,
  next: NavigationGuardNext
) => {
  
  const title = i18n.global.t(`title.${to.name?.toString()}`);

  to.meta = reactive({
    title: title,
  }),
  //@ts-ignore
  window.document.title = `LibrePM | ${title}`;
  next();
};
