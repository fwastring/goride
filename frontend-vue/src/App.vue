<template>
  <div class="wrapper">
    <nav>
      <RouterLink to="/">Home</RouterLink>
      <RouterLink to="/about">About</RouterLink>
    </nav>
    <RouterView />
  </div>
</template>

<script setup>
import HelloWorld from './components/HelloWorld.vue'
</script>

<style scoped>
.wrapper {
  width: 100%;
  display: flex;
  flex-direction: column; /* Stack items vertically */
  align-items: center; /* Center items horizontally */
  justify-content: flex-start; /* Align items to the top */
}

/* Navigation bar styles */
nav {
  width: 100%; /* Full width */
  font-size: 100px; /* Large font size */
  text-align: center; /* Center text */
  margin-bottom: 20px; /* Space below nav */
  margin-top: 0px;
  font-family: 'FiraCode', sans-serif; /* Replace with your preferred font */
  padding: 10px 0; /* Padding for nav */
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* Subtle shadow */
}

nav a {
  display: inline-block; /* Inline block for proper spacing */
  padding: 0 1rem; /* Padding for links */
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0; /* Remove left border for the first link */
}

nav a.router-link-exact-active {
  color: var(--color-text); /* Active link color */
}

nav a.router-link-exact-active:hover {
  background-color: transparent; /* Hover state for active link */
}

/* Media query adjustments */
@media (min-width: 1024px) {
  nav {
    font-size: 2.5rem; /* Adjust font size */
  }
}
</style>

