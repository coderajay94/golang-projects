<template>
  <div class="container my-4">
    <div class="d-flex justify-content-between">
      <input
        v-model="searchCode"
        class="form-control w-50"
        placeholder="Search company code (e.g., HDFCBANK)"
      />
      <button @click="fetchStock" class="btn btn-primary ms-2">Search</button>
    </div>

    <div v-if="stock" class="card mt-4 p-4" style="position: relative">
      <!-- Heart icon to show when added to watchlist -->
      <h4>{{ stock.name }} ({{ stock.companyCode }})</h4>

      <div class="row">
        <div class="col-md-6"><strong>Price:</strong> ₹{{ stock.price }}</div>
        <div class="col-md-6">
          <strong>52W High/Low:</strong> ₹{{ stock.week52High }} / ₹{{
            stock.week52Low
          }}
        </div>
        <div class="col-md-6">
          <strong>Market Cap:</strong> {{ stock.marketCap }}
        </div>
        <div class="col-md-6">
          <strong>P/E Ratio:</strong> {{ stock.peRatio }}
        </div>
        <div class="col-md-6">
          <strong>Day Range:</strong> {{ stock.dayRange }}
        </div>
        <div class="col-md-6">
          <strong>Daily Change:</strong> {{ stock.dailyChange }}
          {{ stock.dailyChangePercentage }}
        </div>
        <div class="col-md-6">
          <strong>Potential Profit %:</strong>
          {{ stock.potentialProfitPercentage.toFixed(2) }}%
        </div>
        <div class="col-md-6">
          <strong>Sector / Industry:</strong> {{ stock.sector }} /
          {{ stock.industry }}
        </div>
        <div class="col-md-6">
          <strong>52 Week High:</strong> {{ stock.week52High }}
        </div>
        <div class="col-md-6">
          <strong>52 Week Low:</strong> {{ stock.week52Low }}
        </div>
      </div>
      <br />
      <br />
      <!-- Shortened Company Profile -->
      <p>{{ shortProfile }}</p>

      <!-- Button to toggle full profile visibility -->
      <button
        v-if="stock && stock.companyProfile"
        class="btn btn-link"
        @click="toggleProfile"
      >
        {{ showFullProfile ? "Show Less" : "Load More" }}
      </button>

      <button
        class="btn mt-4 float-end"
        :class="isInWatchlist ? 'btn-secondary' : 'btn-success'"
        :disabled="isInWatchlist"
        @click="openModal"
      >
        {{ isInWatchlist ? "Already in Watchlist" : "Add to Watchlist" }}
      </button>
      <!-- Remove from Watchlist Button -->
      <button
        v-if="isInWatchlist"
        class="btn btn-outline-danger mt-3"
        @click="removeFromWatchlist"
      >
        ❌ Remove from Watchlist
      </button>
    </div>

    <!-- Modal -->
    <div v-if="showModal" class="modal-overlay">
      <div class="custom-modal">
        <div class="modal-content p-4">
          <h5 class="modal-title">Add to Watchlist</h5>
          <p>
            <strong>{{ stock.name }}</strong> (₹{{ stock.price }})
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
              ₹{{ gainAmount.toFixed(2) }} ({{ gainPercentage.toFixed(2) }}%)
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

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(33, 37, 41, 0.85); /* dark gray with transparency */
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1050;
}

.custom-modal {
  background: white;
  border-radius: 8px;
  max-width: 500px;
  width: 90%;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.5);
}
</style>

<script>
export default {
  data() {
    return {
      searchCode: "",
      stock: null,
      showModal: false,
      targetPrice: "",
      watchlist: [],
      showFullProfile: false, // Flag to toggle full profile view
    };
  },
  computed: {
    gainAmount() {
      // Ensure targetPrice is treated as a number
      const target = parseFloat(this.targetPrice);
      return target ? target - this.stock.price : 0;
    },
    gainPercentage() {
      const target = parseFloat(this.targetPrice);
      return target
        ? ((target - this.stock.price) / this.stock.price) * 100
        : 0;
    },
    isInWatchlist() {
      return (
        this.stock &&
        this.watchlist.some(
          (entry) => entry.companyCode === this.stock.companyCode
        )
      );
    },
    shortProfile() {
      if (this.stock && this.stock.companyProfile) {
        return this.showFullProfile
          ? this.stock.companyProfile
          : this.stock.companyProfile.slice(0, 200) + "...";
      }
      return "";
    },
  },
  mounted() {
    this.checkIfInWatchlist();

    const lastStock = localStorage.getItem("lastSearchedStock");
    if (lastStock) {
      this.stock = JSON.parse(lastStock);
      this.searchCode = this.stock.companyCode;
    }
  },
  methods: {
    removeFromWatchlist() {
      const savedList = JSON.parse(localStorage.getItem("watchlist") || "[]");

      const updatedList = savedList.filter(
        (item) => item.companyCode !== this.stock.companyCode
      );

      localStorage.setItem("watchlist", JSON.stringify(updatedList));
      this.checkIfInWatchlist();

      alert(`${this.stock.name} removed from your watchlist!`);
    },
    fetchStock() {
      fetch(`http://localhost:8080/companies/${this.searchCode}`)
        .then((res) => res.json())
        .then((data) => {
          this.stock = data;
          localStorage.setItem("lastSearchedStock", JSON.stringify(data));
        })
        .catch((err) => console.error(err));
    },
    openModal() {
      this.targetPrice = "";
      this.showModal = true;
    },
    closeModal() {
      this.showModal = false;
    },
    addToWatchlist() {
      const entry = {
        name: this.stock.name,
        companyCode: this.stock.companyCode,
        price: this.stock.price,
        targetPrice: parseFloat(this.targetPrice),
        gainAmount: this.gainAmount,
        gainPercentage: this.gainPercentage,
      };

      // Load existing watchlist
      const savedList = JSON.parse(localStorage.getItem("watchlist") || "[]");

      // Remove if already exists
      const filtered = savedList.filter(
        (item) => item.companyCode !== entry.companyCode
      );

      // Add new one
      const updatedList = [...filtered, entry];

      // Save updated list
      localStorage.setItem("watchlist", JSON.stringify(updatedList));

      this.showModal = false;
      this.checkIfInWatchlist();
      alert(`${entry.name} added to your watchlist!`);
    },
    checkIfInWatchlist() {
      const storedWatchlist =
        JSON.parse(localStorage.getItem("watchlist")) || [];
      this.watchlist = storedWatchlist;
    },
    toggleProfile() {
      this.showFullProfile = !this.showFullProfile;
    },
  },
};
</script>
