// NOTE: 这里负责 注册、登录、用户信息的 api 
// 规范: 导出的函数首字母大写

import { http } from "@/utils/request";
import { RegisterRequest, LoginRequest, ProfileType } from "./profile_types"

export const Register = (data: RegisterRequest) => {
    return http.Post(`/v1/register`, {data})
}

export const Login = (data: LoginRequest) => {
    return http.Post<ResultType<ProfileType>>(`/v1/login`, {data})
}