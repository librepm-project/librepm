import {
    NavigationGuardNext,
    RouteLocationNormalized,
    RouteLocationNormalizedLoaded,
  } from 'vue-router';
import { getToken } from '@/lib/cookie';
import { useSettingStore } from '@/store/setting.store';

export const authorizationMiddleware = async (
    to: RouteLocationNormalized,
    _from: RouteLocationNormalizedLoaded,
    next: NavigationGuardNext
  ) => {
  if (to.name === 'onboard') {
    return next();
  }

  const settingStore = useSettingStore();
  if (!settingStore.settings.length) {
    await settingStore.fetch();
  }
  if (settingStore.getValue('onboarded') === 'false') {
    return next({ name: 'onboard' });
  }

  const accessToken = getToken();
  if (to.name !== 'login' && to.name !== 'register' && !accessToken) {
    next({ name: 'login' });
  } else {
    next();
  }
};
