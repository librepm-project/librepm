import { defineStore } from 'pinia';
import { Tracker } from '@/lib/interfaces/tracker.interface';
import { trackerMock } from '@/lib/mock-data';
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
    getTracker(trackerId: number) {
      this.current = trackerMock[0];
    },

    async getTrackers() {
      this.index = await trackerApi.index()
    },
  },
});
