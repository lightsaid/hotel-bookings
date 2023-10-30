import { http } from "@/utils/request";

// NOTE: 这里负责 酒店、客房、客房类型的 api 
// 规范: 导出的函数首字母大写

export const ListHotels = () => {
    return http.Get<ResultType<string>>(`v1/hotels`, {})
}