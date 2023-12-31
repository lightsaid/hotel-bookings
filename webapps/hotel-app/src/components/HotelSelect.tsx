import { Fragment, useEffect, useState } from "react";
import { Listbox, Transition } from "@headlessui/react";
import { CheckIcon } from "@heroicons/react/20/solid";
import { RiHotelLine } from "react-icons/ri";
import { useBaseDataStore, useHomeStore } from "@/stores"
import { HotelModel } from "@/api";

export function HotelSelect() {
    const { hotels } = useBaseDataStore()
    const { changeHotel }  = useHomeStore()
    const [selected, setSelected] = useState<HotelModel>(hotels[0]);

    useEffect(()=>{
        setSelected(hotels[0])
    },[hotels])

    useEffect(()=>{
        changeHotel(selected)
    },[selected])

    return (
        <div className="w-[220px] h-[50px] box-border">
            <Listbox value={selected} onChange={setSelected}>
                <div className="relative">
                    <Listbox.Button
                        className="
						relative w-full border cursor-pointer border-slate-500 min-h-[50px]
						rounded-lg bg-white py-1 pl-8 pr-2 text-left shadow-md focus:outline-none 
						focus-visible:border-indigo-500 focus-visible:ring-2 focus-visible:ring-white 
						focus-visible:ring-opacity-75 focus-visible:ring-offset-2 focus-visible:ring-offset-orange-300 sm:text-sm"
                    >
                        <div>
                            <div className="absolute left-2 translate-y-[58%]">
                                <RiHotelLine className="text-lg" />
                            </div>
                            <p className="text-[12px] font-bold text-slate-500">
                                酒店
                            </p>
                            <p className="block truncate text-sm text-slate-900">
                                {selected?.title}
                            </p>
                        </div>
                    </Listbox.Button>
                    <Transition
                        as={Fragment}
                        leave="transition ease-in duration-100"
                        leaveFrom="opacity-100"
                        leaveTo="opacity-0"
                    >
                        <Listbox.Options
                            className="
								absolute mt-1 max-h-60 w-full overflow-auto rounded-md
								bg-white py-1 text-base shadow-lg ring-1 ring-black 
								ring-opacity-5 focus:outline-none sm:text-sm"
                        >
                            {hotels.map((h, index) => (
                                <Listbox.Option
                                    key={index}
                                    className={({ active }) =>
                                        `relative cursor-default select-none py-2 pl-10 pr-4 ${
                                            active
                                                ? "bg-sky-100 text-sky-900"
                                                : "text-gray-900"
                                        }`
                                    }
                                    value={h}
                                >
                                    {({ selected }) => (
                                        <>
                                            <span
                                                className={`block truncate ${
                                                    selected
                                                        ? "font-medium"
                                                        : "font-normal"
                                                }`}
                                            >
                                                {h.title}
                                            </span>
                                            {selected ? (
                                                <span className="absolute inset-y-0 left-0 flex items-center pl-3 text-sky-600">
                                                    <CheckIcon
                                                        className="h-5 w-5 "
                                                        aria-hidden="true"
                                                    />
                                                </span>
                                            ) : null}
                                        </>
                                    )}
                                </Listbox.Option>
                            ))}
                        </Listbox.Options>
                    </Transition>
                </div>
            </Listbox>
        </div>
    );
}
