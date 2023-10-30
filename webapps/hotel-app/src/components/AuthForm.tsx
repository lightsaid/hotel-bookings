import { useState, useCallback } from "react";
import { Dialog } from "@headlessui/react";
import { Modal } from "./base/Modal";
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import { Input } from "./base/Input";
import { Button } from "./base/Button";
import { FaWeixin } from "react-icons/fa6";
import { Tab } from "@headlessui/react";

type Variant = "LOGIN" | "REGISTER";
type Titles = {
    [key in Variant]: string;
};
const titles: Titles = {
    LOGIN: "登录，欢迎回来!",
    REGISTER: "注册，欢迎你到来!",
};
export const AuthForm = () => {
    const [variant, setVariant] = useState<Variant>("LOGIN");
    const [isLoading, setIsLoading] = useState(false);

    const toggleVariant = useCallback(() => {
        if (variant === "LOGIN") {
            setVariant("REGISTER");
        } else {
            setVariant("LOGIN");
        }
    }, [variant]);

    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<FieldValues>({
        defaultValues: {
            name: "",
            email: "",
            password: "",
        },
    });

    const onSubmit: SubmitHandler<FieldValues> = (data) => {
        setIsLoading(true);

        if (variant === "REGISTER") {
        }

        if (variant === "LOGIN") {
        }
    };

    const socialAction = (action: string) => {
        setIsLoading(true);
    };

    return (
        <Modal>
            <Dialog.Title
                as="h3"
                className="text-lg font-medium leading-6 text-gray-900"
            ></Dialog.Title>

            <Tab.Group>
                <Tab.List className="flex space-x-1 rounded-xl bg-blue-900/20 p-1">
                    <Tab>账号登录</Tab>
                    <Tab>短信登录</Tab>
                </Tab.List>
                <Tab.Panels className="mt-2">
                    <Tab.Panel className="rounded-xl bg-white p-3 ring-white ring-opacity-60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2">
                        <div className="sm:mx-auto sm:w-full sm:max-w-md">
                            <div className="bg-white px-4 py-8 sm:px-10">
                                <form
                                    className="space-y-6"
                                    onSubmit={handleSubmit(onSubmit)}
                                >
                                    {variant === "REGISTER" && (
                                        <Input
                                            disabled={isLoading}
                                            register={register}
                                            errors={errors}
                                            required
                                            id="name"
                                            label="用户名"
                                        />
                                    )}
                                    <Input
                                        disabled={isLoading}
                                        register={register}
                                        errors={errors}
                                        required
                                        id="phone_number"
                                        label="手机号码"
                                        type="text"
                                    />
                                    <Input
                                        disabled={isLoading}
                                        register={register}
                                        errors={errors}
                                        required
                                        id="password"
                                        label="密码"
                                        type="password"
                                    />
                                    {variant === "REGISTER" && (
                                        <Input
                                            disabled={isLoading}
                                            register={register}
                                            errors={errors}
                                            required
                                            id="repassword"
                                            label="确认密码"
                                            type="password"
                                        />
                                    )}
                                    <div>
                                        <Button
                                            disabled={isLoading}
                                            type="submit"
                                        >
                                            {variant === "LOGIN"
                                                ? "Sign in"
                                                : "Register"}
                                        </Button>
                                    </div>
                                </form>
                            </div>
                        </div>
                    </Tab.Panel>
                    <Tab.Panel className="rounded-xl bg-white p-3 ring-white ring-opacity-60 ring-offset-2 ring-offset-blue-400 focus:outline-none focus:ring-2">
                        <div className="sm:mx-auto sm:w-full sm:max-w-md">
                            <div className="bg-white px-4 py-8 sm:px-10">
                                <form
                                    className="space-y-6"
                                    onSubmit={handleSubmit(onSubmit)}
                                >
                                    {variant === "REGISTER" && (
                                        <Input
                                            disabled={isLoading}
                                            register={register}
                                            errors={errors}
                                            required
                                            id="name"
                                            label="用户名"
                                        />
                                    )}
                                    <Input
                                        disabled={isLoading}
                                        register={register}
                                        errors={errors}
                                        required
                                        id="phone_number"
                                        label="手机号码"
                                        type="text"
                                    />
                                    <Input
                                        disabled={isLoading}
                                        register={register}
                                        errors={errors}
                                        required
                                        id="smscode"
                                        label="短信验证码"
                                        type="text"
                                    />
                                    {variant === "REGISTER" && (
                                        <Input
                                            disabled={isLoading}
                                            register={register}
                                            errors={errors}
                                            required
                                            id="repassword"
                                            label="确认密码"
                                            type="password"
                                        />
                                    )}
                                    <div>
                                        <Button
                                            disabled={isLoading}
                                            type="submit"
                                        >
                                            {variant === "LOGIN"
                                                ? "Sign in"
                                                : "Register"}
                                        </Button>
                                    </div>
                                </form>
                            </div>
                        </div>
                    </Tab.Panel>
                </Tab.Panels>
            </Tab.Group>
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
                <div>
                    {variant === "LOGIN" ? "创建新账号？" : "已经有账号？"}
                </div>
                <div
                    onClick={toggleVariant}
                    className="underline cursor-pointer"
                >
                    {variant === "LOGIN" ? "注册" : "登录"}
                </div>
            </div>
        </Modal>
    );
};
