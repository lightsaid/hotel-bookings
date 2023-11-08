// NOTE: 这里负责 酒店、客房、客房类型的 api
// 规范: 导出的函数首字母大写

import { http } from "@/utils/request";
import { HotelModel, RoomType } from "./hotels_types";

// ListHotels 获取所有酒店
export const ListHotels = (params: { page_num: number; page_size: number }) => {
    return http.Get<ListResultType<HotelModel>>(
        `/v1/hotels`,
        { params: params },
        { error: true, success: false },
    );
};


// ListRoomTypes 获取所有客房类型
export const ListRoomTypes = () => {
    return http.Get<ResultType<RoomType[]>>("/v1/room_types", {}, {error: true, success: false})
}