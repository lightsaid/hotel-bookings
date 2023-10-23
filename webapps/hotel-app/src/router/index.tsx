import { createBrowserRouter } from "react-router-dom";
import { MainLayout } from "../layouts/MainLayout";
import { Bookings } from "../pages/Bookings";
import { Home } from "../pages/Home";
import { Room } from "../pages/Room";

export const router = createBrowserRouter([
    {
        path: "/",
        element: <MainLayout />,
        children: [
            {
                path: "/",
                element: <Home />,
            },
            {
                path: "/room/:id",
                element: <Room />,
            },
            {
                path: "/bookings",
                element: <Bookings />,
            },
        ],
    },
]);
