<template>
  <div class="route-creator-container">
    <h2 class="title">Create a New Route</h2>
    <div class="input-group">
      <InputText v-model="fromValue" placeholder="Start Address" class="input-field" />
      <InputText v-model="toValue" placeholder="End Address" class="input-field" />
    </div>
    <Button 
      type="button" 
      label="Create" 
      icon="pi pi-plus" 
      @click="addRoute" 
      :disabled="!canSubmit"
      class="create-button"
    />
  </div>
</template>

<script>
import { ref, computed } from 'vue';
import { useRouteStore } from '@/stores/routeStore';

export default {
  setup() {
    const routeStore = useRouteStore();

    // Local state for start and end addresses
    const fromValue = ref('');
    const toValue = ref('');

    // Computed property to check if both fields are filled
    const canSubmit = computed(() => fromValue.value && toValue.value);

    // Method to handle the route creation
    const addRoute = async () => {
      const newRoute = {
        from: fromValue.value,
        to: toValue.value,
      };
		routeStore.addRoute(newRoute)

        fromValue.value = '';
        toValue.value = '';
    };

    return {
      fromValue,
      toValue,
      canSubmit,
      addRoute,
    };
  }
};
</script>

<style scoped>
.route-creator-container {
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

.create-button {
  margin-top: 10px; /* Space above the button */
  width: 100%; /* Full width button */
}
</style>

