import { Fragment, useEffect, useState } from "react";
import { Listbox, Transition } from "@headlessui/react";
import { CheckIcon } from "@heroicons/react/20/solid";
// import { CgSortAz, CgSortZa } from "react-icons/cg";
import { useHomeStore, priceSortList } from "@/stores";


export function PriceSort() {
    const { priceSort, changePriceSort } = useHomeStore()
    const [selected, setSelected] = useState(priceSort);

    useEffect(()=>{
        changePriceSort(selected)
    }, [selected])

    return (
        <div className="w-[160px] h-[50px] box-border">
            <Listbox value={selected} onChange={setSelected}>
                <div className="relative">
                    <Listbox.Button
                        className="
						relative w-full border cursor-pointer border-slate-500 
						rounded-lg bg-white py-1 pl-8 pr-2 text-left shadow-md focus:outline-none 
						focus-visible:border-indigo-500 focus-visible:ring-2 focus-visible:ring-white 
						focus-visible:ring-opacity-75 focus-visible:ring-offset-2 focus-visible:ring-offset-orange-300 sm:text-sm"
                    >
                        <div>
                            <div className="absolute left-1 translate-y-[40%]">
                                <selected.icon className=" text-2xl" />
                            </div>
                            <p className="text-[12px] font-bold text-slate-500">
                                排序
                            </p>
                            <p className="block truncate text-sm text-slate-900">
                                {selected.name}
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
                            bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm"
                        >
                            {priceSortList.map((item, index) => (
                                <Listbox.Option
                                    key={index}
                                    className={({ active }) =>
                                        `relative cursor-default select-none py-2 pl-7 pr-3 ${
                                            active
                                                ? "bg-sky-100 text-sky-900"
                                                : "text-gray-900"
                                        }`
                                    }
                                    value={item}
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
                                                {item.name}
                                            </span>
                                            {selected ? (
                                                <span className="absolute inset-y-0 left-0 flex items-center pl-2 text-sky-600">
                                                    <CheckIcon
                                                        className="h-4 w-4 "
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
