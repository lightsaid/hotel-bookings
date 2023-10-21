import { useState } from "react";
// import reactLogo from "./assets/react.svg";
// import viteLogo from "/vite.svg";
import "./App.css";

import { Button } from "./components/Button";

function App() {
    const [count, setCount] = useState(0);

    return (
        <>
            <h1 className="text-3xl font-bold underline text-red-600">Hello world!</h1>
			<Button onClick={()=>alert(1)} variant="danger">Login</Button>
			<Button onClick={()=>alert(1)} variant="secondary" size={"lg"} className="bg-red-500 text-white hover:bg-red-800">Login</Button>
			<Button onClick={()=>alert(1)} variant="primary" size={"md"} className="bg-red-500 text-white hover:bg-red-800">Login</Button>

        </>
    );
}

export default App;
