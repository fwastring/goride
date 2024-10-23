<template>
  <div class="route-searcher-container">
    <h2 class="title">Search for a Route</h2>
    <div class="input-group">
      <InputText v-model="fromValue" placeholder="Start Address" class="input-field" />
      <InputText v-model="toValue" placeholder="End Address" class="input-field" />
    </div>
    <Button 
      type="button" 
      label="Search" 
      icon="pi pi-search" 
      @click="findRoute" 
      :disabled="!canSubmit"
      class="search-button"
    />
  </div>
</template>

<script>
import { ref, computed } from 'vue';
import { useRouteStore } from '@/stores/routeStore';
import { searchRoute } from '@/services/api'; // Import the API function to search routes

export default {
  setup() {
    const routeStore = useRouteStore();

    // Local state for start and end addresses
    const fromValue = ref('');
    const toValue = ref('');

    // Computed property to check if both fields are filled
    const canSubmit = computed(() => fromValue.value && toValue.value);

    // Method to handle the route search
    const findRoute = async () => {
      const searchParams = {
        from: fromValue.value,
        to: toValue.value,
      };

      try {
        // Call the API to search for the route
        const response = await searchRoute(searchParams);
        const foundRouteIds = response.data; // Assuming the API returns a list of routes

        // Update the Pinia store with the found routes
        routeStore.filterRoutes(foundRouteIds);
		routeStore.setRiderAddresses(fromValue.value, toValue.value)

        // Clear the input fields
      } catch (error) {
        console.error('Error searching for routes:', error);
      }
    };

    return {
      fromValue,
      toValue,
      canSubmit,
      findRoute,
    };
  }
};
</script>

<style scoped>
.route-searcher-container {
  max-width: 400px; /* Set a max width for the container */
  margin: 0 auto; /* Center the container */
  padding: 20px; /* Add padding */
  border: 1px solid #ccc; /* Light border */
  border-radius: 10px; /* Rounded corners */
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1); /* Subtle shadow */
}

.title {
  font-size: 24px; /* Title size */
  margin-bottom: 20px; /* Space below title */
  text-align: center; /* Center title */
}

.input-group {
  display: flex; /* Flexbox for horizontal layout */
  flex-direction: column; /* Stack inputs vertically */
  gap: 10px; /* Space between inputs */
}

.input-field {
  padding: 10px; /* Input padding */
  border: 1px solid #ccc; /* Border for inputs */
  border-radius: 5px; /* Rounded corners */
}

.search-button {
  margin-top: 10px; /* Space above the button */
  width: 100%; /* Full width button */
}
</style>

