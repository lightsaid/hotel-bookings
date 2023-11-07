import { useState, useCallback, Fragment, memo } from "react";
import { Dialog, Tab } from "@headlessui/react";
import { Modal } from "./base/Modal";
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import { Input } from "./base/Input";
import { Button } from "./base/Button";
import { FaWeixin } from "react-icons/fa6";
import { twMerge } from "tailwind-merge";
import { useNavigate } from "react-router-dom";
import { useProfileStore } from "@/stores";
import { LoginType } from "@/api";
import { toast } from "react-hot-toast";
import { useModalContext } from "@/context/ModalContext";
const tabs = [
    {
        key: 0,
        title: "密码登录",
    },
    {
        key: 1,
        title: "短信登录",
    },
];

export const AuthForm = () => {
    const { closeModal } = useModalContext();
    const { login } = useProfileStore();
    const [isLoading, setIsLoading] = useState(false);
    const [currentTab, setCurrentTab] = useState(0);

    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<FieldValues>({
        defaultValues: {
            phone_number: "",
            password: "",
            sms_code: "",
        },
    });

    const onSubmit: SubmitHandler<FieldValues> = (data) => {
        console.log(currentTab);
        console.log(data);
        let arg = {
            login_type: currentTab as LoginType,
            sms_code: data.sms_code,
            password: data.password,
            phone_number: data.phone_number,
        };
        login(arg).then((res) => {
            toast.success("登录成功");
            closeModal();
        });
    };

    return (
        <Modal>
            <Dialog.Title
                as="h3"
                className="text-lg font-medium leading-6 text-gray-900"
            ></Dialog.Title>

            <Tab.Group
                selectedIndex={currentTab}
                onChange={(index) => setCurrentTab(index)}
            >
                <Tab.List className="flex space-x-8 justify-center">
                    {tabs.map((item) => (
                        <Tab key={item.key} as={Fragment}>
                            {({ selected }) => (
                                <button
                                    className={twMerge(
                                        "text-gray-500 border-b-[3px] text-xl font-medium rounded-b-sm border-transparent",
                                        selected
                                            ? " border-sky-500  text-gray-900"
                                            : "",
                                    )}
                                >
                                    {item.title}
                                </button>
                            )}
                        </Tab>
                    ))}
                </Tab.List>

                {tabs.map((item) => (
                    <Tab.Panels className="mt-2" key={item.key}>
                        <Tab.Panel className="rounded-xl">
                            <div className="sm:mx-auto sm:w-full sm:max-w-md">
                                <div className="bg-white px-4 py-8 sm:px-10">
                                    <form
                                        className="space-y-6"
                                        onSubmit={handleSubmit(onSubmit)}
                                    >
                                        <Input
                                            disabled={isLoading}
                                            register={register}
                                            errors={errors}
                                            required
                                            id="phone_number"
                                            label="手机号码"
                                            type="text"
                                        />
                                        {currentTab === 0 ? (
                                            <Input
                                                disabled={isLoading}
                                                register={register}
                                                errors={errors}
                                                required
                                                id="password"
                                                label="密码"
                                                type="password"
                                            />
                                        ) : (
                                            <div className="flex items-end">
                                                <div className="flex-[3]">
                                                    <Input
                                                        disabled={isLoading}
                                                        register={register}
                                                        errors={errors}
                                                        required
                                                        id="sms_code"
                                                        label="短信验证码"
                                                        type="text"
                                                    />
                                                </div>
                                                <button className="flex-[2] bg-sky-500 h-[36px] -ml-1 rounded-r-md text-white hover:bg-sky-600">
                                                    获取验证码
                                                </button>
                                            </div>
                                        )}

                                        <div>
                                            <Button
                                                disabled={isLoading}
                                                type="submit"
                                                className="w-full flex justify-center mt-8 text-lg"
                                            >
                                                登录
                                            </Button>
                                        </div>
                                    </form>
                                </div>
                            </div>
                        </Tab.Panel>
                    </Tab.Panels>
                ))}
            </Tab.Group>

            <Footer />
        </Modal>
    );
};

const Footer = memo(() => {
    const navigate = useNavigate();
    const gotoRegister = () => {
        navigate("/register");
    };
    return (
        <Fragment>
            <div className="mt-6">
                <div className="relative">
                    <div className="absolute inset-0 flex items-center">
                        <div className="w-full border-t border-gray-300" />
                    </div>
                    <div className="relative flex justify-center text-sm">
                        <span className="bg-white px-2 text-gray-500">
                            其他方式登录
                        </span>
                    </div>
                </div>

                <div className="mt-6 flex justify-center gap-2">
                    <FaWeixin className="text-4xl text-green-500 cursor-pointer" />
                </div>
            </div>
            <div className="flex gap-1 justify-center text-sm mt-6 px-2 text-gray-500">
                <div>创建新账号？</div>
                <div
                    onClick={gotoRegister}
                    className="underline cursor-pointer"
                >
                    注册
                </div>
            </div>
        </Fragment>
    );
});
