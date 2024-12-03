import {
  NavigationGuardNext,
  RouteLocationNormalized,
  RouteLocationNormalizedLoaded,
} from 'vue-router';

export const pageTitleMiddleware = (
  to: RouteLocationNormalized,
  _from: RouteLocationNormalizedLoaded,
  next: NavigationGuardNext
) => {
  //@ts-ignore
  window.document.title = `LibrePM | ${to.meta.title}`;
  next();
};
