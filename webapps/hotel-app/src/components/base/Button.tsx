import { ButtonHTMLAttributes, ReactNode } from "react";
import { twMerge } from "tailwind-merge";
import { cva, VariantProps } from "class-variance-authority";
import { clsx } from "clsx";

// 继承原有的属性、方法
interface ButtonProps
    extends ButtonHTMLAttributes<HTMLButtonElement>,
        VariantProps<typeof buttonVariants> {
    children: ReactNode;
}

export const Button = ({
    children,
    className,
    variant,
    size,
    ...props
}: ButtonProps) => {
    return (
        <button
            {...props}
            className={twMerge(
                clsx(buttonVariants({ variant, size }), className),
            )}
        >
            {children}
        </button>
    );
};

// 第一个参数是通用的，第二个参数配置变体
const buttonVariants = cva("rounded-md", {
    variants: {
        // 变体的 key 不能跟原有的属性同名，如这里就不能使用type，因为button也有type属性
        variant: {
            primary:
                "flex-none rounded-md font-semibold hover:bg-sky-600 bg-sky-500 text-white border-[1px] dark:shadow-highlight/20",
            // TODO: 待设计
            secondary:
                "flex-none rounded-md font-semibold hover:bg-sky-600 bg-sky-500 text-white border-[1px] dark:shadow-highlight/20",
            danger: "flex-none rounded-md font-semibold hover:bg-red-600 bg-red-500 text-white border-[1px] dark:shadow-highlight/20",
        },
        size: {
            sm: "leading-6 py-1 px-1 text-sm  py-1 shadow-sm",
            md: "leading-6 py-1.5 px-3 shadow-md text-base",
            lg: "text-xl px-4 py-2",
        },
    },
    defaultVariants: {
        variant: "primary",
        size: "md",
    },
});
