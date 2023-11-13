// 管理酒店基础数据如（客房类型、预定状态、支付状态等）
import { create } from "zustand"
import { persist, createJSONStorage } from "zustand/middleware"
import { HotelModel, ListHotels, ListRoomTypes, RoomType } from "@/api"

type State = {
    hotels: HotelModel[]
    roomTypes: RoomType[]
}

type Actions = {
    getRoomTypes: () => Promise<ResultType<RoomType[]>>
    getHotels: () => Promise<ListResultType<HotelModel>>
}

export const useBaseDataStore = create<State & Actions>()(
    persist(
        ((set, state) => ({
            hotels: [],
            roomTypes: [],
            getRoomTypes: async () => {
                const res = await ListRoomTypes()
                set({...state, roomTypes: res.result})
                return res
            },
            getHotels: async () => {
                const res = await ListHotels({page_num: 1, page_size: 100})
                set({...state, hotels: res.list})
                return res
            }
        }))
    , {
        name: "base_data",
        storage: createJSONStorage(()=>sessionStorage),
    })
)