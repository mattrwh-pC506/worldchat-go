import {useEffect} from "react";
import {
    createHashRouter,
    useNavigate,
} from "react-router-dom";

import App from './App';
import Login from './Login';
import Chat from './Chat';
import {isAuthorized} from "../auth";

export const router = createHashRouter([
    {
        path: "/",
        element: <App />,
    },
    {
        path: "/login",
        element: <Login />,
    },
    {
        path: "/chat",
        element: <Chat />,
    },
]);

export const useAuthorizedRouting = () => {
    const navigate = useNavigate();

    useEffect(() => {
        isAuthorized().then((authorized) => {
            if (authorized) {
                navigate("/chat")
            } else {
                navigate("/login")
            }
        })
    }, [])
}