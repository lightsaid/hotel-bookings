import { Link } from "react-router-dom";
import { Button } from "./base/Button";
import { BiUserCircle } from "react-icons/bi";
import { AuthForm } from "./AuthForm";
import { useModalContext, ModalContextProvider } from "../context/ModalContext";
import { useEffect } from "react";
import { AvatarMenu } from "./AvatarMenu";

const Container = () => {
    const { setConfig, openModal } = useModalContext();
    useEffect(() => {
        setConfig({ panelClass: "max-w-[540px]" });
    }, []);
    return (
        <>
            <section className="sticky top-0 z-10 border-b-[1px] bg-white shadow-md border-x-slate-400">
                <header className="relative flex flex-wrap justify-between items-center w-[1024px] max-w-[1024px] mx-auto py-2">
                    <div>
                        <Link
                            to="/"
                            className=" text-2xl font-bold text-zinc-800"
                        >
                            Hotel Bookings
                        </Link>
                    </div>
                    <div className="text-right ml-auto">
                        <Button
                            variant={"primary"}
                            size="md"
                            className="flex justify-items-center items-center ml-auto"
                            onClick={() => openModal()}
                        >
                            <BiUserCircle className="h-[1.35rem] w-6 inline-block" />
                            <span className="ml-1">登录</span>
                        </Button>
                    </div>
                    <div className="ml-6">
                        <AvatarMenu />
                    </div>
                </header>
            </section>
            <AuthForm />
        </>
    );
};

export const NavHeader = () => {
    return (
        <ModalContextProvider>
            <Container />
        </ModalContextProvider>
    );
};
