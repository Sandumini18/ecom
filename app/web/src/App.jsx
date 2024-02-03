import React from "react";
import "./App.css";
import { CartProvider } from "react-use-cart";

// rq
import { QueryClient, QueryClientProvider } from "react-query";
import { ReactQueryDevtools } from "react-query/devtools";

// routes
import { ApplicationRouter } from "./Routes";

const queryClient = new QueryClient();

const App = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <ReactQueryDevtools initialIsOpen={false} />
      <CartProvider
        id="jamie"
        onItemAdd={(item) => {
          alert(`Item ${item.id} added!`);
        }}
        onItemUpdate={(item) => alert(`Item ${item.id} updated.!`)}
        onItemRemove={() => alert(`Item removed!`)}
      >
        <ApplicationRouter />
      </CartProvider>
    </QueryClientProvider>
  );
};

export default App;
