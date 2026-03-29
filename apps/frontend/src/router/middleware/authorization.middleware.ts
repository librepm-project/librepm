import {
    NavigationGuardNext,
    RouteLocationNormalized,
    RouteLocationNormalizedLoaded,
  } from 'vue-router';
import { getToken } from '@/lib/cookie';

export const authorizationMiddleware = (
    to: RouteLocationNormalized,
    _from: RouteLocationNormalizedLoaded,
    next: NavigationGuardNext
  ) => {
  const accessToken = getToken();
  if (to.name !== 'login' && to.name !== 'register' && !accessToken) {
    next({ name: 'login' });
  } else {
    next();
  }
};
