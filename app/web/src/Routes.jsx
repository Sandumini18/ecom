// react router
import { useRoutes } from "react-router-dom";
// pages
import Home from "./pages/Home";
import { ItemDetails } from "./pages/ItemDetails";
import { Dashboard } from "./pages/Dashboard";
import { CheckOut } from "./pages/checkout";
import { Categories } from "./pages/Categories";
import { CategorieDetails } from "./pages/CategorieDetails";

export const ApplicationRouter = () => {
  let routes = useRoutes([
    {
      path: "/",
      element: <Home />,
    },
    {
      path: "/item/:slug",
      element: <ItemDetails />,
    },
    {
      path: "/dashboard",
      element: <Dashboard />,
    },
    {
      path: "/categories",
      element: <Categories />,
    },
    {
      path: "/category/:slug",
      element: <CategorieDetails />,
    },

    {
      path: "/checkout",
      element: <CheckOut />,
    },
  ]);

  return routes;
};
