import { Fragment } from "react";
import { Popover, Transition } from "@headlessui/react";
import { LuUsers2 } from "react-icons/lu";
import { Button } from "./base/Button";
import { AiOutlineMinus, AiOutlinePlus } from "react-icons/ai";


export function CustomFilter() {
    return (
        <div className=" top-16 w-full max-w-sm px-4">
            <Popover className="relative">
                {({ open }) => (
                    <>
                        <Popover.Button
                            className={
								`w-[220px] h-[50px] relative border cursor-pointer border-slate-500  box-border
								rounded-lg bg-white shadow-md  py-1 pl-8 pr-2 text-left
								focus:outline-none focus-visible:border-indigo-500 focus-visible:ring-2 focus-visible:ring-white 
								sm:text-sm`}
                        >
                            <div>
                            <div className="absolute left-2 translate-y-[58%]">
                                <LuUsers2 className="text-xl" />
                            </div>
                            <p className="text-[12px] font-bold text-slate-500">
                                旅客
                            </p>
                            <p className="block truncate text-sm text-slate-900">
                                {`1位旅客, 1间客房`}
                            </p>
                        </div>
                        </Popover.Button>
                        <Transition
                            as={Fragment}
                            enter="transition ease-out duration-200"
                            enterFrom="opacity-0 translate-y-1"
                            enterTo="opacity-100 translate-y-0"
                            leave="transition ease-in duration-150"
                            leaveFrom="opacity-100 translate-y-0"
                            leaveTo="opacity-0 translate-y-1"
                        >
                            <Popover.Panel className="absolute w-full left-0 z-10 mt-3  transform ">
                                <ul className="bg-white rounded-lg border border-slate-100 shadow-lg px-4 pb-3 lg:max-w-3xl">
                                    <li className="py-3">
										<span className="text-lg text-slate-800">客房1</span>
										<div className="flex justify-between mt-3">
											<span className="text-base text-slate-600">旅客</span>
											<div> 
												<Button size="sm" className="bg-transparent hover:bg-transparent hover:border-sky-500 px-2 py-1 text-sm"><AiOutlineMinus className=" text-slate-600" /></Button>
													<span className="text-base text-slate-600 px-2">2</span>
												<Button size="sm" className="bg-transparent hover:bg-transparen hover:border-sky-500  px-2 py-1 text-sm"><AiOutlinePlus className=" text-slate-600"  /></Button></div>
										</div>
									</li>
									<li>
										<span className="text-lg text-slate-800">客房2</span>
										<div className="flex justify-between mt-3">
											<span className="text-base text-slate-600">旅客</span>
											<div> 
												<Button size="sm" className="bg-transparent hover:bg-transparent hover:border-sky-500 px-2 py-1 text-sm"><AiOutlineMinus className=" text-slate-600" /></Button>
													<span className="text-base text-slate-600 px-2">2</span>
												<Button size="sm" className="bg-transparent hover:bg-transparen hover:border-sky-500  px-2 py-1 text-sm"><AiOutlinePlus className=" text-slate-600"  /></Button></div>
										</div>
									</li>

									<div className="mt-4 mb-3 text-right"><Button size="sm" className="px-2 py-1.5">添加一间客房</Button></div>
                                </ul>
                            </Popover.Panel>
                        </Transition>
                    </>
                )}
            </Popover>
        </div>
    );
}

