import { HotelModel, RoomType } from "@/api";
import { create } from "zustand";
import * as dayjs from 'dayjs'
import { CgSortAz, CgSortZa } from "react-icons/cg";

const dateLayout = "YYYY-MM-DD"

function setDate(d: number): string {
    return dayjs(new Date()).add(d, 'day').format(dateLayout)
}

const date = {startDate: setDate(1), endDate: setDate(2)}
export type HoteDate = typeof date

export type PriceSortType = {
    name: string;
    icon: any;
    value: string;
}

export const priceSortList: PriceSortType[] = [
    { name: "价格：低到高", icon: CgSortZa, value: "ASC" },
    { name: "价格：高到低", icon: CgSortAz, value: "DESC" },
];

type State = {
    hotel: HotelModel | null
    date: HoteDate
    roomType: RoomType | null
    priceSort: PriceSortType
}

type Actions = {
    changeHotel: (h: HotelModel) => void
    changeDate: (d: HoteDate) => void
    changeRoomType: (roomType: RoomType) => void
    changePriceSort: (value: PriceSortType) => void
}

export const useHomeStore = create<State & Actions>()(
    (set) => ({
        hotel: null,
        date: date,
        roomType: null,
        priceSort: priceSortList[0],
        changeHotel: (h: HotelModel) => {
            set((state) => ({ ...state, hotel: h }))
        },
        changeDate: (d: HoteDate) => {
            set((state) => ({...state, date: d}))
        },
        changeRoomType: (rt: RoomType) => {
            set((state) => ({...state, roomType: rt}))
        },
        changePriceSort: (value: PriceSortType) => {
            set((state) => ({...state, priceSort: value}))
        }
    })
)