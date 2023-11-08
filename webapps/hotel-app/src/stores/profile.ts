import { create } from "zustand";
import { createJSONStorage, persist } from "zustand/middleware";

import { ProfileType, LoginRequest, Login } from "@/api";

type State = {
    profile: ProfileType | null;
};

type Actions = {
    login: (params: LoginRequest) => Promise<ResultType<ProfileType>>;
};

export const useProfileStore = create<State & Actions>()(persist((set) => ({
    profile: null,
    login: async (params: LoginRequest) => {
        const res = await Login(params);
        set({ profile: res.result });
        return res;
    },
}),{
    name: "profile_store", // persist 中间件默认是保存在localStorage
    // storage: createJSONStorage(()=>sessionStorage) // 存储在 sessionStorage 
    // partialize: (state) => ({profile: state.profile}) // 默认是保存所有状态，如果想保存某一个，直接partialize函数返回
}));
