import { Outlet } from "react-router-dom";
import { NavHeader } from "../components/NavHeader";
export const MainLayout = () => {
    return (
        <section>
            <NavHeader />
            <Outlet />
        </section>
    );
};
