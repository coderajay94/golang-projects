<template>
  <div class="container my-5">
    <h2 class="text-center mb-4">Explore Stocks</h2>

    <div class="table-responsive">
      <table class="table table-bordered table-striped">
        <thead>
          <tr>
            <th @click="sortBy('name')" class="sortable text-start">
              Stock Name <i :class="sortClass('name')"></i>
            </th>
            <th @click="sortBy('price')" class="sortable">
              Price <i :class="sortClass('price')"></i>
            </th>
            <th @click="sortBy('dayRange')" class="sortable">
              Today's Range <i :class="sortClass('dayRange')"></i>
            </th>
            <th @click="sortBy('week52High')" class="sortable">
              52W High <i :class="sortClass('week52High')"></i>
            </th>
            <th class="sortable text-nowrap text-center">
              52W Low <i :class="sortClass('week52Low')"></i>
            </th>

            <th @click="sortBy('potentialGain')" class="sortable text-nowrap">
              Potential Gain <i :class="sortClass('potentialGain')"></i>
            </th>
            <th
              @click="sortBy('dailyChangePercentage')"
              class="sortable text-nowrap"
            >
              Daily % Change <i :class="sortClass('dailyChangePercentage')"></i>
            </th>
            <th>Watchlist</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="stock in sortedStocks" :key="stock.companyCode">
            <td class="text-start">
              <strong>{{ stock.name }} </strong>
            </td>
            <td>₹{{ stock.price }}</td>
            <td>{{ stock.dayRange }}</td>
            <td>₹{{ stock.week52High }}</td>
            <td>₹{{ stock.week52Low }}</td>
            <td class="text-success">
              {{ stock.potentialProfitPercentage.toFixed(2) }}%
            </td>
            <td class="text-success">{{ stock.dailyChangePercentage }}%</td>
            <td>
              <input
                type="checkbox"
                :checked="isInWatchlist(stock.companyCode)"
                @change="toggleWatchlist(stock)"
              />
              <span
                v-if="isInWatchlist(stock.companyCode)"
                style="color: black"
              >
                Added
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Modal -->
    <div v-if="showModal && selectedStock" class="modal-overlay">
      <div class="custom-modal">
        <div class="modal-content p-4">
          <h5 class="modal-title">Add to Watchlist</h5>
          <p>
            <strong>{{ selectedStock.name }}</strong> (₹{{
              selectedStock.price
            }})
          </p>
          <input
            v-model="targetPrice"
            class="form-control"
            placeholder="Enter target price"
            type="number"
            min="0"
          />
          <div v-if="targetPrice" class="mt-2">
            <p>
              <strong>Potential Gain:</strong>
              ₹{{ (targetPrice - selectedStock.price).toFixed(2) }} (
              {{
                getPotentialGainPercentage(selectedStock, targetPrice).toFixed(
                  2
                )
              }}% )
            </p>
          </div>

          <div class="d-flex justify-content-end mt-3">
            <button class="btn btn-secondary me-2" @click="closeModal">
              Cancel
            </button>
            <button class="btn btn-primary" @click="addToWatchlist">
              Save
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Stocks",
  data() {
    return {
      name: "",
      stocks: [],
      watchlist: [],
      showModal: false,
      selectedStock: null,
      targetPrice: "",
      currentSort: "name",
      currentSortDir: "asc",
    };
  },
  mounted() {
    this.fetchStocks();
    this.loadWatchlist();
  },
  computed: {
    sortedStocks() {
      return this.stocks.slice().sort((a, b) => {
        const aValue = a[this.currentSort];
        const bValue = b[this.currentSort];

        if (aValue > bValue) return this.currentSortDir === "asc" ? 1 : -1;
        if (aValue < bValue) return this.currentSortDir === "asc" ? -1 : 1;
        return 0;
      });
    },
  },
  methods: {
    fetchStocks() {
      fetch("http://localhost:8080/companies")
        .then((res) => res.json())
        .then((data) => {
          this.stocks = data;
        })
        .catch((err) => console.error("Error fetching stocks:", err));
    },
    loadWatchlist() {
      this.watchlist = JSON.parse(localStorage.getItem("watchlist") || "[]");
    },
    isInWatchlist(companyCode) {
      return this.watchlist.some((item) => item.companyCode === companyCode);
    },
    toggleWatchlist(stock) {
      if (this.isInWatchlist(stock.companyCode)) {
        this.removeFromWatchlist(stock);
      } else {
        this.openModal(stock);
      }
    },
    openModal(stock) {
      this.selectedStock = stock;
      this.targetPrice = "";
      this.showModal = true;
    },
    closeModal() {
      this.showModal = false;
      this.selectedStock = null;
    },
    addToWatchlist() {
      if (!this.selectedStock) return;

      const entry = {
        name: this.selectedStock.name,
        companyCode: this.selectedStock.companyCode,
        price: this.selectedStock.price,
        targetPrice: parseFloat(this.targetPrice),
        gainAmount: this.targetPrice - this.selectedStock.price,
        gainPercentage: this.getPotentialGainPercentage(
          this.selectedStock,
          this.targetPrice
        ),
      };

      this.watchlist.push(entry);
      localStorage.setItem("watchlist", JSON.stringify(this.watchlist));
      this.closeModal();
      if (this.selectedStock?.name) {
        alert(`${this.selectedStock.name} added to your watchlist!`);
      } else {
        alert("Stock added to your watchlist!");
      }
    },
    removeFromWatchlist(stock) {
      this.watchlist = this.watchlist.filter(
        (item) => item.companyCode !== stock.companyCode
      );
      localStorage.setItem("watchlist", JSON.stringify(this.watchlist));
    },
    getPotentialGain(stock) {
      return ((stock.week52High - stock.price) / stock.price) * 100;
    },
    getPotentialGainPercentage(stock, targetPrice) {
      return ((targetPrice - stock.price) / stock.price) * 100;
    },
    sortBy(column) {
      if (this.currentSort === column) {
        this.currentSortDir = this.currentSortDir === "asc" ? "desc" : "asc";
      } else {
        this.currentSort = column;
        this.currentSortDir = "asc";
      }
    },
    sortClass(column) {
      if (this.currentSort !== column) return "text-muted";
      return this.currentSortDir === "asc" ? "text-primary" : "text-danger";
    },
  },
};
</script>

<style scoped>
.table-responsive {
  margin-top: 20px;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th,
td {
  text-align: center;
  padding: 12px;
}

th {
  background-color: #007bff;
  color: white;
  cursor: pointer;
}

th.sortable:hover {
  background-color: #0056b3;
}

tr:nth-child(even) {
  background-color: #f2f2f2;
}

tr:hover {
  background-color: #d9edf7;
}

input[type="checkbox"] {
  cursor: pointer;
}

.modal-backdrop {
  background-color: rgba(0, 0, 0, 0.7);
}

.modal-dialog {
  max-width: 600px;
}

.modal-content {
  border-radius: 8px;
  background-color: white;
  padding: 30px;
}

.modal-title {
  font-size: 1.25rem;
  font-weight: 600;
}

.modal-body {
  padding: 1rem;
}

.btn-primary {
  background-color: #007bff;
  border-color: #007bff;
}

.btn-secondary {
  background-color: #6c757d;
  border-color: #6c757d;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1050;
}

.custom-modal {
  background: white;
  border-radius: 10px;
  max-width: 600px;
  width: 100%;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
}
th {
  white-space: nowrap;
}
</style>
