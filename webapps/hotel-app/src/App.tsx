import * as React from "react";
import { RouterProvider } from "react-router-dom";
import { router } from "./router";
import { Toaster } from "react-hot-toast";

function App() {
    return (
        <React.Fragment>
            <Toaster />
            <RouterProvider router={router} />
        </React.Fragment>
    );
}

export default App;
