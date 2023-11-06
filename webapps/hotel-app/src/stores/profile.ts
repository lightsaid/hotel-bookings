import { create } from "zustand";
import { ProfileType, LoginRequest, Login } from "@/api";

type State = {
    profile: ProfileType | null;
};

type Actions = {
    login: (params: LoginRequest) => Promise<ResultType<ProfileType>>;
};

export const useProfileStore = create<State & Actions>()((set) => ({
    profile: null,
    login: async (params: LoginRequest) => {
        const res = await Login(params);
        set({ profile: res.result });
        return res;
    },
}));
