import { defineStore } from 'pinia';
import { Release } from '@/lib/interfaces/release.interface';
import releaseApi from '@/api/release.api';
import { createCrudActions } from './utils/crudMixin';

interface ReleaseStore {
  current: Release | null;
  index: Release[];
}

export const useReleaseStore = defineStore('release', {
  state: (): ReleaseStore => {
    return {
      current: null,
      index: [],
    };
  },
  actions: {
    ...createCrudActions<Release>(releaseApi),
  },
});
