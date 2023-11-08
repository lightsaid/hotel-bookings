export type HotelModel = {
    id: number;
    title: string;
    code: string;
    address: string;
    created_at: string;
    updated_at: string;
    is_deleted?: boolean;
};

export type RoomType = {
    "id": number;
    "room_type": string;
    "room_label": string;
};
