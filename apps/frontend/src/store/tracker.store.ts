import { defineStore } from 'pinia';
import { Tracker } from '@/lib/interfaces/tracker.interface';
import trackerApi from '@/api/tracker.api';
import { createCrudActions } from './utils/crudMixin';

interface TrackerStore {
  current: Tracker | null;
  index: Tracker[];
}

export const useTrackerStore = defineStore('tracker', {
  state: (): TrackerStore => {
    return {
      current: null,
      index: [],
    };
  },
  actions: {
    ...createCrudActions<Tracker>(trackerApi),
  },
});
