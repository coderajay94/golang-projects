<template>
  <div class="container mt-4">
    <h2>Watchlist</h2>
    <!-- Table placeholder for now -->
    <table class="table">
      <thead>
        <tr>
          <th>Company</th>
          <th>Current Price</th>
          <th>Target Price</th>
          <th>Potential Profit</th>
        </tr>
      </thead>
      <tbody>
        <!-- Loop for displaying future watchlist items -->
        <tr v-for="stock in watchlist" :key="stock.companyCode">
          <td>{{ stock.name }}</td>
          <td>{{ stock.price }}</td>
          <td>{{ stock.targetPrice || "-" }}</td>
          <td>
            {{ calculateProfit(stock) }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: "StockWatchlist",
  data() {
    return {
      watchlist: [], // Initially empty
    };
  },
  mounted() {
    // In real use, you'd fetch this from localStorage or API
    try {
      const saved = localStorage.getItem("watchlist");
      this.watchlist = saved ? JSON.parse(saved) : [];
    } catch (e) {
      console.error("Failed to load watchlist:", e);
    }
  },
  methods: {
    calculateProfit(stock) {
      if (!stock.targetPrice) return "-";
      const diff = parseFloat(stock.targetPrice) - parseFloat(stock.price);
      const percent = ((diff / parseFloat(stock.price)) * 100).toFixed(2);
      return `${diff.toFixed(2)} (${percent}%)`;
    },
  },
};
</script>
