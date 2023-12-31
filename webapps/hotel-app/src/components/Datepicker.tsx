import { default as TwDatepicker } from "react-tailwindcss-datepicker";
import { CgCalendarDates } from "react-icons/cg";
import { useHomeStore} from "@/stores"

const Datepicker = () => {
    const { date, changeDate } = useHomeStore();

    return (
        <li className="ml-4 w-[225px] h-[50px] box-border border shadow-md rounded-lg border-slate-500 cursor-default bg-white">
            <div className="relative w-full z-50  bg-white pl-8 py-1  pr-3 text-left rounded-lg">
                <div className="absolute left-2 translate-y-[50%]">
                    {" "}
                    <CgCalendarDates className="text-lg" />
                </div>
                <p className="text-[12px] font-bold text-slate-500">日期</p>
                <p className="block min-h-[20px] truncate text-sm text-slate-900">
                    {date?.startDate?.toLocaleString()}{" "}
                    {date.startDate == null ? "" : "~"}{" "}
                    {date?.endDate?.toLocaleString()}
                </p>
                <div className="w-full h-full absolute top-0 left-0 bottom-0 right-0">
                    <TwDatepicker
                        i18n={"zh"}
                        value={date}
                        onChange={changeDate}
                        showShortcuts={false}
                        inputClassName="home-datepicker h-[50px] top-0 left-0 bottom-0 right-0 opacity-0 border-none cursor-pointer"
                    />
                </div>
            </div>
        </li>
    );
};

export default Datepicker;
