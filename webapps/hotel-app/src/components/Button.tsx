import { ButtonHTMLAttributes, ReactNode } from "react";
import { twMerge } from "tailwind-merge";
import { cva, VariantProps } from "class-variance-authority";
import { clsx } from "clsx"

// 继承原有的属性、方法
interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement>, VariantProps<typeof buttonVariants> {
    children: ReactNode;
}

export const Button = ({ children, className, variant, size, ...props }: ButtonProps) => {
    return (
        <button
            {...props}
			className={ twMerge(clsx(buttonVariants({variant, size}))) }
            // className={twMerge(
            //     "border border-blue-200 px-2 rounded-sm",
            //     className,
            // )}
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
            primary: "border-2 border-black px-2 hover:bg-neutral-200",
            secondary: "border-2 border-white px-2 bg-cyan-500",
            danger: "border-none text-white bg-red-500",
        },
        size: {
            sm: "text-sm px-1 py-0",
            md: "text-base  px-2 py-1",
            lg: "text-xl px-4 py-2",
        },
    },
	defaultVariants: {
		variant: "primary",
		size: "md"
	}
});
