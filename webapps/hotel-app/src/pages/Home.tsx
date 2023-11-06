import { HotelSelect } from "../components/HotelSelect";
import { PriceSort } from "../components/PriceSort";
import AvatarImage from "../assets/default_avatar.png";
import { useState, useEffect } from "react";
import { Button } from "../components/base/Button";
import Datepicker from "@/components/Datepicker";
// import { CustomFilter } from "../components/CustomFilter";
import RoomTypeSelect from "@/components/RoomTypeSelect";
import { ListHotels, Login, LoginRequest } from "@/api";
import { useCounterStore } from "@/stores/counter";

export const Home = () => {
    const { count, increment, decrement, getHotels, hotels } =
        useCounterStore();
    const query = {
        page_num: 1,
        page_size: 10,
    };
    const loginRequest: LoginRequest = {
        phone_number: "13876543212",
        password: "abc123",
        sms_code: "1111",
        login_type: 0,
    };
    useEffect(() => {
        // TODO: 测试接口
        // ListHotels(query).then(res=>{})
        // Login(loginRequest).then(res=>{})
        (async () => {
            let data = await ListHotels(query);
            // console.log(data)
        })();

        getHotels().then((res) => {
            console.log(res);
        });
        console.log("hotels: ", hotels);
    }, []);

    return (
        <section className="max-w-[1024px] mx-auto">
            {/* 搜索栏 */}
            <ul className="relative flex justify-start mt-5 mb-5">
                <li>
                    <HotelSelect />
                </li>
                <Datepicker />
                {/* TODO: 后续确定在放开 */}
                {/* <li className="ml-4">
                    <CustomFilter />
                </li> */}
                <li className="ml-4">
                    <RoomTypeSelect />
                </li>

                <li className="ml-4">
                    <PriceSort />
                </li>
                <li className="ml-4">
                    <Button
                        variant="primary"
                        size={"lg"}
                        className="px-7 h-[50px]"
                    >
                        搜 索
                    </Button>
                </li>
            </ul>

            {/* Rooms */}
            <div className="grid sm:grid-cols-2 md:grid-cols-3 gap-x-8 gap-y-8">
                {/* TODO: 测试 */}
                <div>
                    <span className="text-lg font-bold text-cyan-500">
                        {count}
                    </span>
                    <Button size="md" onClick={() => increment(3)}>
                        increment
                    </Button>
                    <Button size="md" onClick={() => decrement(1)}>
                        increment
                    </Button>
                </div>

                <div className="item">
                    <img
                        className="w-full h-48 object-cover rounded-t-lg"
                        src={AvatarImage}
                        alt=""
                    />
                    <div>
                        <h2 className="font-bold text-2xl">Blog title 1</h2>
                        <p>
                            Lorem ipsum, dolor sit amet consectetur adipisicing
                            elit. Nemo assumenda porro inventore repellendus
                            ipsum.
                        </p>
                        <a href="#">Read more</a>
                    </div>
                </div>
                <div className="item">
                    <img
                        className="w-full h-48 object-cover rounded-t-lg"
                        src={AvatarImage}
                        alt=""
                    />
                    <div>
                        <h2 className="font-bold text-2xl">Blog title 2</h2>
                        <p>
                            Lorem ipsum, dolor sit amet consectetur adipisicing
                            elit. Nemo assumenda porro inventore repellendus
                            ipsum.
                        </p>
                        <a href="#">Read more</a>
                    </div>
                </div>
                <div className="item">
                    <img
                        className="w-full h-48 object-cover rounded-t-lg"
                        src={AvatarImage}
                        alt=""
                    />
                    <div className="desc">
                        <h2 className="font-bold text-2xl">Blog title 3</h2>
                        <p>
                            Lorem ipsum, dolor sit amet consectetur adipisicing
                            elit. Nemo assumenda porro inventore repellendus
                            ipsum.
                        </p>
                        <a href="#">Read more</a>
                    </div>
                </div>
                <div className="item">
                    <img
                        className="w-full h-48 object-cover rounded-t-lg"
                        src={AvatarImage}
                        alt=""
                    />
                    <div className="desc">
                        <h2 className="font-bold text-2xl">Blog title 4</h2>
                        <p>
                            Lorem ipsum, dolor sit amet consectetur adipisicing
                            elit. Nemo assumenda porro inventore repellendus
                            ipsum.
                        </p>
                        <a href="#">Read more</a>
                    </div>
                </div>
                <div className="item">
                    <img
                        className=" w-full h-48 object-cover rounded-t-lg"
                        src={AvatarImage}
                        alt=""
                    />
                    <div className="desc">
                        <h2 className="font-bold text-2xl">Blog title 5</h2>
                        <p>
                            Lorem ipsum, dolor sit amet consectetur adipisicing
                            elit. Nemo assumenda porro inventore repellendus
                            ipsum.
                        </p>
                        <a href="#">Read more</a>
                    </div>
                </div>
                <div className="item">
                    <img
                        className="w-full h-48 object-cover rounded-t-lg"
                        src={AvatarImage}
                        alt=""
                    />
                    <div className="desc">
                        <h2 className="font-bold text-2xl">Blog title 6</h2>
                        <p>
                            Lorem ipsum, dolor sit amet consectetur adipisicing
                            elit. Nemo assumenda porro inventore repellendus
                            ipsum.
                        </p>
                        <a href="#">Read more</a>
                    </div>
                </div>
            </div>
        </section>
    );
};
