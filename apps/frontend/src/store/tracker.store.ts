import { defineStore } from 'pinia';
import { Tracker } from '@/lib/interfaces/tracker.interface';
import trackerApi from '@/api/tracker.api';


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
    async getTracker(trackerId: string) {
      this.current = await trackerApi.show(trackerId);
    },

    async getTrackers() {
      this.index = await trackerApi.index()
    },

    async postTracker(tracker: Omit<Tracker, 'id'>) {
      const newTracker = await trackerApi.create(tracker);
      this.index.push(newTracker);
      return newTracker;
    },

    async putTracker(trackerId: string, tracker: Omit<Tracker, 'id'>) {
      const updatedTracker = await trackerApi.update(trackerId, tracker);
      const index = this.index.findIndex(t => t.id === trackerId);
      if (index !== -1) {
        this.index[index] = updatedTracker;
      }
      return updatedTracker;
    },

    async deleteTracker(trackerId: string) {
      await trackerApi.destroy(trackerId);
      this.index = this.index.filter(t => t.id !== trackerId);
    },
  },
});
