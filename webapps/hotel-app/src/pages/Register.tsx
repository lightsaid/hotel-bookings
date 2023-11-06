import { useState, useCallback } from "react";
import { Dialog } from "@headlessui/react";
import { Modal } from "@/components/base/Modal";
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import { Input } from "@/components/base/Input";
import { Button } from "@/components/base/Button";
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
const Register = () => {
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

    return (
        <div className="sm:mx-auto sm:w-full sm:max-w-md">
            <div className="bg-white px-4 py-8 sm:px-10">
                <form className="space-y-6" onSubmit={handleSubmit(onSubmit)}>
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
                        <Button disabled={isLoading} type="submit">
                            {variant === "LOGIN" ? "Sign in" : "Register"}
                        </Button>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default Register;
