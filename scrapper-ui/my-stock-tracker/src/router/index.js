import { createRouter, createWebHistory } from "vue-router";
import StockDashboard from "../components/StockDashboard.vue";
import StockWatchlist from "../components/StockWatchlist.vue";
import Stocks from "../components/Stocks.vue";

const routes = [
  { path: "/", redirect: "/dashboard" },
  { path: "/dashboard", component: StockDashboard },
  { path: "/watchlist", component: StockWatchlist },
  { path: "/stocks", component: Stocks },

  // You can add /stocks later if needed
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
