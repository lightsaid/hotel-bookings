import { createContext, useContext } from "react";

export const contextFactory = <T extends unknown | null>() => {
    // 创建一个context
    const context = createContext<T | undefined>(undefined);

    // 通过闭包的形式缓存起 context
    const useCtx = () => {
        const ctx = useContext(context);
        if (ctx === undefined) {
            // <context.Provider value={}> </context.Provider> value 必须要有值
            throw new Error("useContext 必须在带有值的Provider中。");
        }

        return ctx;
    };

    // 最后把父类要使用 context 和子类使用的 useCtx 方法返出去
    return [context, useCtx] as const;
};
