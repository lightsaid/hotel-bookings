import { ReactNode, useState, useCallback } from "react";
import { contextFactory } from "./helpers/contextFactory";

type ModalConfig = {
    isOpen: boolean;
    panelClass: string; // 设置 Dialog.Panel 元素的 class，如何设置宽：max-w-[800px]
    onClose: (callback?: () => void) => void; // 点击模态框阴影层或者按 Esc 触发的回调。
};

export type ModalContextType = {
    config: ModalConfig;
    setConfig: (cfg: Partial<ModalConfig>) => void;
    openModal: () => void;
    closeModal: (callback?: () => void) => void;
};

const [ModalContext, useModalContext] = contextFactory<ModalContextType>();

const ModalContextProvider = ({ children }: { children: ReactNode }) => {
    const [config, newConfig] = useState<ModalConfig>({
        isOpen: false,
        panelClass: "",
        onClose: () => {
            closeModal();
        },
    });

    const setConfig = (cfg: Partial<ModalConfig>) => {
        newConfig((prev) => ({ ...prev, ...cfg }));
    };

    const openModal = useCallback(() => {
        newConfig((prev) => ({ ...prev, isOpen: true }));
    }, []);

    const closeModal = useCallback((callback?: () => void) => {
        newConfig((prev) => ({ ...prev, isOpen: false }));
        callback && callback();
    }, []);

    const value = { config, setConfig, openModal, closeModal };

    return (
        <ModalContext.Provider value={value}>{children}</ModalContext.Provider>
    );
};

export { useModalContext, ModalContextProvider };
