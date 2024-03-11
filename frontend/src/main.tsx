import React from 'react'
import ReactDOM from 'react-dom/client'
import "./tailwind.css"
import {
  createHashRouter,
  RouterProvider,
} from "react-router-dom";
import { router } from './router';

const hashRouter=createHashRouter(router)

ReactDOM.createRoot(document.getElementById('root')!).render(
  <RouterProvider router={hashRouter} />
)
