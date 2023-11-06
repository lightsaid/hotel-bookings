// NOTE: 这里负责 酒店、客房、客房类型的 api
// 规范: 导出的函数首字母大写

import { http } from "@/utils/request";
import { HotelType } from "./hotels_types";

export const ListHotels = (params: { page_num: number; page_size: number }) => {
    return http.Get<ListResultType<HotelType>>(
        `/v1/hotels`,
        { params: params },
        { error: true, success: false },
    );
};
