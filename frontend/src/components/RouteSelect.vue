<template>
  <div class="route-container">
    <div v-for="route in routes" :key="route.id" class="route-card">
      <Card background-color="#ff66ff" style="width: 25rem; overflow: hidden">
        <template #title>{{ route.from }} - {{ route.to }}</template>
        <template #footer>
          <div class="flex gap-4 mt-1">
            <Button label="Delete" severity="secondary" outlined class="w-full" @click="handleCancel(route.id)" />
            <Button label="Display" class="w-full" @click="displayRoute(route)" />
            <Button label="Join" severity="info" class="w-full" @click="joinRouteHandler(route)" />
          </div>
        </template>
      </Card>
    </div>
  </div>
</template>

<script>
import { onMounted, computed } from 'vue';
import { useRouteStore } from '@/stores/routeStore'; // Import the store
import Button from 'primevue/button';
import Card from 'primevue/card';
import { joinRoute } from '@/services/api'; // Import the API function to add routes

export default {
  components: {
    Button,
    Card,
  },
  setup(props, { emit }) {
    const routeStore = useRouteStore(); // Use the route store
  const routes = computed(() => routeStore.routes);

    const fetchRoute = async () => {
      await routeStore.fetchRoutes(); // Call the store action to fetch routes
		console.log(routeStore.routes)
    };

    const handleCancel = async (routeId) => {
      try {
		  routeStore.removeRoute(routeId)
      } catch (error) {
        console.error('Error removing route:', error);
      }
    };

    const joinRouteHandler = async (route) => {
      try {
		  console.log(route)
		  joinRoute({
			  "trip_id": route.id,
			  "rider_id": 12,
			  "from": routeStore.rider_start_address,
			  "to": routeStore.rider_end_address,
		  })
      } catch (error) {
        console.error('Error removing route:', error);
      }
    };

    const displayRoute = (route) => {
      emit('display-route', route); // Emit event to parent to handle route display
    };

    // onMounted(() => {
    //   fetchRoute(); // Fetch routes when the component is mounted
    // });

    return {
      routes, // Return the store to access routes in the template
	  joinRouteHandler,
      fetchRoute,
      handleCancel,
      displayRoute,
    };
  },
};
</script>

<style scoped>
/* Add any styles if needed */
</style>

