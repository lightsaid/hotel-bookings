import { ref, computed } from "vue";
import { defineStore } from "pinia";

// export const useCounterStore = defineStore("counter", () => {
//     const count = ref(0);
//     const doubleCount = computed(() => count.value * 2);
//     function increment() {
//         count.value++;
//     }

//     return { count, doubleCount, increment };
// });


export const minSidebarWidth = "64px"
export const maxSidebarWidth = "210px"

export type AppStore = {
    isCollapse: boolean;
    sidebarWidth: string;  // 210px | 64px
}

export const useAppStore = defineStore("app", {
    state: () => ({isCollapse: false, sidebarWidth: maxSidebarWidth} as AppStore),
    getters: {},
    actions: {
        setMenuStatus() {
            this.isCollapse = !this.isCollapse
            if (this.sidebarWidth === maxSidebarWidth) {
                this.sidebarWidth = minSidebarWidth
            }else{
                this.sidebarWidth = maxSidebarWidth
            }
        }
    },
})