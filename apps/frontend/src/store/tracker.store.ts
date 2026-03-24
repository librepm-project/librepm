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

    // Alias methods for backward compatibility
    async getTracker(trackerId: string) {
      return this.getCurrentItem(trackerId);
    },

    async getTrackers() {
      return this.getAllItems();
    },

    async postTracker(tracker: Omit<Tracker, 'id'>) {
      return this.createItem(tracker);
    },

    async putTracker(trackerId: string, tracker: Omit<Tracker, 'id'>) {
      return this.updateItem(trackerId, tracker);
    },

    async deleteTracker(trackerId: string) {
      return this.deleteItem(trackerId);
    },
  },
});
