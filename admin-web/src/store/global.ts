import { defineStore } from "pinia";

export const useGlobalStore = defineStore("global", {
  state: () => {
    return {
      count: 0,
    };
  },

  getters: {
    getCount: (state) => state.count * 2,
  },

  actions: {
    updateCount() {
      this.count += 1;
    },
  },
});
