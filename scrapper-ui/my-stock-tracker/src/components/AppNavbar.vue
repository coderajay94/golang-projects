<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-white shadow-sm">
    <div class="container-fluid">
      <a class="navbar-brand fw-bold fs-4 text-primary" href="#"
        >📈 Stock Tracker</a
      >

      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarNav"
        aria-controls="navbarNav"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>

      <div
        class="collapse navbar-collapse justify-content-center"
        id="navbarNav"
      >
        <ul class="navbar-nav gap-4">
          <li class="nav-item">
            <router-link to="/dashboard" class="nav-link fw-semibold"
              >Dashboard</router-link
            >
          </li>
          <li class="nav-item">
            <router-link to="/stocks" class="nav-link fw-semibold"
              >Stocks</router-link
            >
          </li>
          <li class="nav-item">
            <router-link to="/watchlist" class="nav-link fw-semibold"
              >Watchlist</router-link
            >
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: "AppNavbar",
};
</script>

<style scoped>
.nav-link {
  font-size: 1.1rem;
  transition: color 0.3s;
}
.nav-link:hover {
  color: #0d6efd;
  text-decoration: underline;
}
</style>
