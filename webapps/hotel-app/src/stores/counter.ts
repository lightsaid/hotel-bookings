import { create } from "zustand";
import { HotelType, ListHotels } from "@/api";

type State = {
    count: number;
    hotels: HotelType[];
};

type Actions = {
    increment: (qty: number) => void;
    decrement: (qty: number) => void;
    getHotels: () => Promise<HotelType[]>;
};

export const useCounterStore = create<State & Actions>()((set) => ({
    count: 0,
    hotels: [],
    increment: (qty: number) => set((state) => ({ count: state.count + qty })),
    decrement: (qty: number) => set((state) => ({ count: state.count - qty })),
    getHotels: async () => {
        const res = await ListHotels({ page_num: 1, page_size: 10 });
        // set(state => ({...state, hotels: res.list}))
        return res.list;
    },
}));
