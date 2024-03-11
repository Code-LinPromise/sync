import Home from "../pages/Home"
import Photo from "../pages/Photo"
import Text from "../pages/Text"
import File from "../pages/File"
import { Navigate } from "react-router-dom"

export const router=[
    {
        path: "/",
        element:(<Navigate to="/text"/>)
    },
    {
        path: "/",
        element: <Home />,
        children: [
            {
                path:'/text',
                element: <Text />,
            },
            {
                path:'/photo',
                element: <Photo />,
            },
            {
                path:'/file',
                element: <File />,
            },
        ],
    },
]