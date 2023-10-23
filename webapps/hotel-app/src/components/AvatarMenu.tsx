import { memo, useState } from "react";
import AvatarImage from "../assets/default_avatar.png";
import { twJoin } from "tailwind-merge"
import { AiOutlineBars, AiOutlineLogin, AiOutlineForm  } from "react-icons/ai"

export const AvatarMenu = memo(() => {
	const [showLink, setShowLink] = useState(false)
    return (
        <div className="relative">
            <button className="
				flex justify-center items-center px-3 py-1 rounded-md border
				border-slate-100 shadow-sm hover:border hover:border-slate-300"

				onClick={()=>setShowLink(prev=>!prev)}
			>
                <img
                    src={AvatarImage}
                    alt="avatar"
                    className="w-[36px] rounded-full bg-slate-300 shadow-sm border border-slate-200 mr-2"
                />
                <span>Lightsaid</span>
            </button>
			<ul className={twJoin(showLink?`block`:`hidden`, `absolute mt-2 border bg-white w-[140px] rounded-md`)}>
				<li onClick={()=>setShowLink(false)} className="px-3 py-2 border-b flex items-center text-slate-600 cursor-pointer hover:bg-slate-100">
					<AiOutlineBars size={19} className="text-slate-600 mr-2" />  我的预订
				</li>
				<li onClick={()=>setShowLink(false)} className="px-3 py-2 border-b flex items-center text-slate-600 cursor-pointer hover:bg-slate-100">
					<AiOutlineForm size={19} className="text-slate-600 mr-2" />  修改资料
				</li>
				<li onClick={()=>setShowLink(false)} className="px-3 py-2 border-b flex items-center text-slate-600 cursor-pointer hover:bg-slate-100">
					<AiOutlineLogin size={19} className="text-slate-600 mr-2" />  退出登录
				</li>
			</ul>
        </div>
    );
});
