import { HotelSelect } from  "../components/HotelSelect";
import { RoomSort } from "../components/RoomSort"
import Datepicker from "react-tailwindcss-datepicker"; 
import AvatarImage from "../assets/default_avatar.png";
import { useState } from "react";
import { Button } from "../components/base/Button";
import { CgCalendarDates } from "react-icons/cg"
import { CustomFilter } from "../components/CustomFilter"

import * as dayjs from 'dayjs'

export const Home = () => {
    
    const [value, setValue] = useState({
        startDate: dayjs().format("YYYY-MM-DD"),
        endDate: dayjs().format("YYYY-MM-DD")
    });

    type IVuew = typeof value 

    const handleValueChange = (newValue: IVuew) => {
        console.log("newValue:", newValue);
        setValue(newValue);
    };
    return (
        <section className="max-w-[1024px] mx-auto">
        <ul className="relative flex justify-start mt-5 mb-5">
            <li>
                <HotelSelect />
            </li>
            <li className="ml-4 w-[225px] h-[50px] box-border border shadow-md rounded-lg border-slate-500 cursor-default bg-white">
                <div className="relative w-full z-50  bg-white pl-8 py-1  pr-3 text-left rounded-lg">
                    <div className="absolute left-2 translate-y-[50%]"> <CgCalendarDates className="text-lg" /></div>
                    <p className="text-[12px] font-bold text-slate-500">日期</p>
                    <p className="block min-h-[20px] truncate text-sm text-slate-900">
                        {value?.startDate?.toLocaleString()} { value.startDate == null? '' : '~' } {value?.endDate?.toLocaleString()}
                    </p>
                    <div className="w-full h-full absolute top-0 left-0 bottom-0 right-0">
                        <Datepicker 
                            i18n={"zh"} 
                            value={value} 
                            onChange={handleValueChange} 
                            showShortcuts={false}
                            inputClassName="home-datepicker h-[50px] top-0 left-0 bottom-0 right-0 opacity-0 border-none cursor-pointer"
                        />
                    </div>
                </div>
            </li>
            <li className="ml-4">
                <CustomFilter />
            </li>
            <li className="ml-4">
                <RoomSort />
            </li>
            <li className="ml-4"><Button variant="primary" size={"lg"} className="px-7">搜 索</Button></li>
        </ul>
        <div className="grid sm:grid-cols-2 md:grid-cols-3 gap-x-8 gap-y-8">
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
                        elit. Nemo assumenda porro inventore repellendus ipsum.
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
                        elit. Nemo assumenda porro inventore repellendus ipsum.
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
                        elit. Nemo assumenda porro inventore repellendus ipsum.
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
                        elit. Nemo assumenda porro inventore repellendus ipsum.
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
                        elit. Nemo assumenda porro inventore repellendus ipsum.
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
                        elit. Nemo assumenda porro inventore repellendus ipsum.
                    </p>
                    <a href="#">Read more</a>
                </div>
            </div>
        </div>
        </section>
    );
};
