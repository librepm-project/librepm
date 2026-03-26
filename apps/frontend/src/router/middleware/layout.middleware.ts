import { NavigationGuardNext, RouteLocationNormalized } from 'vue-router';
import { useLayoutStore } from '@/store/layout.store';

export const layoutMiddleware = (
  to: RouteLocationNormalized,
  _from: RouteLocationNormalized,
  next: NavigationGuardNext
) => {
  const layoutStore = useLayoutStore();
  
  if (to.meta?.hideLayout) {
    layoutStore.setHideLayout(true);
  } else {
    layoutStore.setHideLayout(false);
  }
  
  next();
};
