import {
    NavigationGuardNext,
    RouteLocationNormalized,
    RouteLocationNormalizedLoaded,
  } from 'vue-router';

export const authorizationMiddleware = (
    to: RouteLocationNormalized,
    _from: RouteLocationNormalizedLoaded,
    next: NavigationGuardNext
  ) => {
  const accessToken = localStorage.getItem('accessToken');
  const requiresAuth = to.matched.some((record) => record.meta.authRequired);
  if (requiresAuth && !accessToken) {
    next('login');
  } else {
    next();
  }
};
