// 管理酒店基础数据如（客房类型、预定状态、支付状态等）
import { create } from "zustand"
import { persist, createJSONStorage } from "zustand/middleware"
import { ListRoomTypes, RoomType } from "@/api"

let baseTimer: number | undefined  = undefined 

type State = {
    roomTypes: RoomType[]
}

type Actions = {
    getRoomTypes: () => Promise<ResultType<RoomType[]>>
}

export const useBaseDataStore = create<State & Actions>()(
    persist(
        ((set, state) => ({
            roomTypes: [],
            getRoomTypes: async () => {
                const res = await ListRoomTypes()
                set({...state, roomTypes: res.result})
                // 简单做一下缓存, 5分钟
                baseTimer = setTimeout(()=> {
                    set({...state, roomTypes: []})
                    clearTimeout(baseTimer)
                }, 1000 * 60 * 5)
                return res
            }
        }))
    , {
        name: "base_data",
        storage: createJSONStorage(()=>sessionStorage),
    })
)